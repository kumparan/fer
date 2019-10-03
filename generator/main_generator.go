package generator

import (
	"bytes"
	"fmt"
	"github.com/kumparan/fer/util"
	"io/ioutil"
	"log"

	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	scaffold = "scaffold.sh"
	proto2go = "proto2go.sh"
	template = "template.sh"
	bash     = "bash"
)

type (
	//Generator define generator
	Generator interface {
		Run(serviceName, protoPath string)
	}
	generator struct {
		service   Service
		client    Client
		protofile string
	}
)

//NewGenerator :nodoc:
func NewGenerator() Generator {
	return &generator{}
}

//Generate to generate service
func (g generator) Run(serviceName string, protoPath string) {
	fmt.Println("Creating ", serviceName)
	fmt.Println(serviceName, "Scaffolding ...")
	_ = os.RemoveAll(serviceName)
	_ = os.Mkdir(serviceName, os.ModePerm)
	_ = os.Mkdir(serviceName+"/client", os.ModePerm)

	g.GetTemplates(serviceName)
	serviceURL := "gitlab.kumparan.com/yowez/" + serviceName

	if protoPath != "" {
		protoPath = strings.ReplaceAll(protoPath, "\n", "")
		g.GenerateProto2Go(protoPath)
		protoPath = strings.ReplaceAll(protoPath, "proto", "pb.go")
		fmt.Println(">>> " + serviceName + " <<<")
		fmt.Println("RPC Path : ", protoPath)
		g.protofile = protoPath
		_ = os.Mkdir(serviceName+"/"+g.getProtoFolder(protoPath), os.ModePerm)
		_ = util.CopyFolder(g.getProtoFolder(protoPath)+"/", serviceName+"/"+protoPath+"/")
		g.client = NewRPCClientGenerator(g.protofile, serviceName, serviceURL)
		g.service = NewServiceGenerator(serviceName, serviceURL, g.protofile)
		fmt.Println(serviceName, "Generating client ...")
		g.client.Generate()
		time.Sleep(1500 * time.Millisecond)
		fmt.Println(serviceName, "Generating service&test ...")
		g.service.Generate()
		time.Sleep(1500 * time.Millisecond)
	}

	g.CreateScaffoldScript()
	g.RunScaffold(serviceName)
	fmt.Println(serviceName, "Created")
}

//CreateScaffoldScript to create bash script for scaffolding
func (g generator) CreateScaffoldScript() {
	contents := `#!/usr/bin/env bash
servicename=$1;
find $servicename -type f -exec sed -i '' "s/skeleton-service/$servicename/g" {} \;
cp $servicename/*.example $servicename/config.yml;
cd $servicename;
go mod tidy;
go get;
`
	bt := []byte(contents)
	_ = ioutil.WriteFile("scaffold.sh", bt, 0644)

}

//GetTemplates to get service template
func (g generator) GetTemplates(serviceName string) {
	contents := `#!/usr/bin/env bash
	servicename=$1;
	cd $servicename;
	git init;
	git remote add origin git@gitlab.kumparan.com:yowez/skeleton-service.git;
	git remote -v;
	git fetch;	
	git pull origin master;
	rm -rf .git;
	cd ..;
`
	bt := []byte(contents)
	_ = ioutil.WriteFile(template, bt, 0644)
	defer func() {
		_ = os.Remove(template)
	}()
	cmd := exec.Command(bash, template, serviceName)
	g.runCmd(cmd)

}

//GenerateProto2Go :nodoc:
func (g generator) GenerateProto2Go(path string) {
	contents := `#!/usr/bin/env bash
path=$1;	
pathproto="${path}/*.proto";
pathpbgo="${path}/*.pb.go";
echo $pathproto;
echo $pathpbgo;
protoc --go_out=plugins=grpc:. $pathproto;
ls $pathpbgo | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}';
`
	bt := []byte(contents)
	_ = ioutil.WriteFile(proto2go, bt, 0644)
	defer func() {
		_ = os.Remove(proto2go)
	}()
	pathArr := strings.Split(path, "/")
	pathArr = pathArr[:len(pathArr)-1]
	path = strings.Join(pathArr, "/")
	cmd := exec.Command(bash, proto2go, path)
	g.runCmd(cmd)
}

//RunScaffold to run scaffold script
func (g generator) RunScaffold(serviceName string) {
	cmd := exec.Command(bash, scaffold, serviceName)
	defer func() {
		_ = os.Remove(scaffold)
	}()
	g.runCmd(cmd)
}

func (g generator) runCmd(cmd *exec.Cmd) {
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (g generator) getProtoFolder(path string) string {
	pathArr := strings.Split(path, "/")
	pathArr = pathArr[:len(pathArr)-1]
	path = strings.Join(pathArr, "/")
	return path
}
