package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kumparan/fer/util"

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
		service     Service
		client      Client
		protoFile   string
		serviceName string
	}
)

//NewGenerator :nodoc:
func NewGenerator() Generator {
	return &generator{}
}

//Generate to generate service
func (g generator) Run(serviceName string, protoPath string) {
	g.serviceName = serviceName
	fmt.Println(">>> " + serviceName + " <<<")
	fmt.Println("Creating ", serviceName)
	fmt.Println(serviceName, "Scaffolding ...")
	err := os.Mkdir(serviceName, os.ModePerm)
	if err != nil {
		log.Fatal("folder " + serviceName + " already exist")
	}
	err = os.Mkdir(serviceName+"/client", os.ModePerm)
	if err != nil {
		g.rollbackWhenError("fail create client folder inside " + serviceName)
	}

	g.GetTemplates(serviceName)
	serviceURL := "gitlab.kumparan.com/yowez/" + serviceName

	if protoPath != "" {
		_, err := os.Stat(protoPath)
		if os.IsNotExist(err) {
			g.rollbackWhenError("proto not exists or invalid path")
		}
		protoPath = strings.ReplaceAll(protoPath, "\n", "")
		g.GenerateProto2Go(protoPath)
		protoPath = strings.ReplaceAll(protoPath, "proto", "pb.go")
		fmt.Println("RPC Path : ", protoPath)
		g.protoFile = protoPath
		err = os.Mkdir(serviceName+"/"+g.getProtoFolder(protoPath), os.ModePerm)
		if err != nil {
			g.rollbackWhenError("fail to create dir " + serviceName + "/" + g.getProtoFolder(protoPath))
		}
		err = util.CopyFolder(g.getProtoFolder(protoPath)+"/", serviceName+"/"+protoPath+"/")
		if err != nil {
			g.rollbackWhenError("fail to create dir " + g.getProtoFolder(protoPath) + "/" + serviceName + "/" + protoPath + "/")
		}
		g.client = NewRPCClientGenerator(g.protoFile, serviceName, serviceURL)
		g.service = NewServiceGenerator(serviceName, serviceURL, g.protoFile)
		fmt.Println(serviceName, "Generating client ...")
		err = g.client.Generate()
		if err != nil {
			g.rollbackWhenError("fail generate client " + err.Error())
		}
		time.Sleep(1500 * time.Millisecond)
		fmt.Println(serviceName, "Generating service&test ...")
		err = g.service.Generate()
		if err != nil {
			g.rollbackWhenError("fail generate service&test " + err.Error())
		}
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
`
	bt := []byte(contents)
	err := ioutil.WriteFile(scaffold, bt, 0644)
	if err != nil {
		g.rollbackWhenError("fail when create scaffold script")
	}
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
	err := ioutil.WriteFile(template, bt, 0644)
	if err != nil {
		g.rollbackWhenError("fail when create scaffold script")
	}
	defer func() {
		_ = os.Remove(template)
	}()
	cmd := exec.Command(bash, template, serviceName)
	err = g.runCmd(cmd)
	if err != nil {
		log.Fatal("fail run scaffold script")
	}

}

//GenerateProto2Go :nodoc:
func (g generator) GenerateProto2Go(path string) {
	contents := `#!/usr/bin/env bash
path=$1;	
pathproto="${path}/*.proto";
pathpbgo="${path}/*.pb.go";
protoc --go_out=plugins=grpc:. $pathproto;
ls $pathpbgo | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}';
`
	bt := []byte(contents)
	err := ioutil.WriteFile(proto2go, bt, 0644)
	if err != nil {
		g.rollbackWhenError("fail when run proto script")
	}
	defer func() {
		_ = os.Remove(proto2go)
	}()
	pathArr := strings.Split(path, "/")
	pathArr = pathArr[:len(pathArr)-1]
	path = strings.Join(pathArr, "/")
	cmd := exec.Command(bash, proto2go, path)
	err = g.runCmd(cmd)
	if err != nil {
		g.rollbackWhenError("fail when run proto script")
	}
}

//RunScaffold to run scaffold script
func (g generator) RunScaffold(serviceName string) {
	cmd := exec.Command(bash, scaffold, serviceName)
	defer func() {
		_ = os.Remove(scaffold)
	}()
	err := g.runCmd(cmd)
	if err != nil {
		g.rollbackWhenError("fail on scaffolding")
	}
}

func (g generator) runCmd(cmd *exec.Cmd) error {
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (g generator) getProtoFolder(path string) string {
	pathArr := strings.Split(path, "/")
	pathArr = pathArr[:len(pathArr)-1]
	path = strings.Join(pathArr, "/")
	return path
}

func (g generator) rollbackWhenError(message string) {
	_ = os.RemoveAll(g.serviceName)
	_ = os.Remove(scaffold)
	_ = os.Remove(template)
	_ = os.Remove(proto2go)
	log.Fatal(message)
}
