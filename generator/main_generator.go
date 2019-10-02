package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/kumparan/fer/util"

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
	bash = "bash"
)

type (
	Generator interface {
		Run(serviceName, protoPath string)
	}
	generator struct {
		service Service
		client  Client
	}
)

func NewGenerator() Generator {
	return &generator{}
}

//Generate to generate service
func (g generator) Run(serviceName string, protoPath string) {
	protoPath = strings.ReplaceAll(protoPath, "\n", "")
	GenerateProto2Go(protoPath)
	protoPath = strings.ReplaceAll(protoPath, "proto", "pb.go")
	fmt.Println(">>> " + serviceName + " <<<")
	fmt.Println("RPC Path : ", protoPath)
	protoCoreFile := protoPath
	serviceURL := "gitlab.kumparan.com/yowez/" + serviceName

	fmt.Println("Creating ", serviceName)

	_ = os.RemoveAll(serviceName)
	_ = os.Mkdir(serviceName, os.ModePerm)
	GetTemplates(serviceName)
	_ = os.Mkdir(serviceName+"/service", os.ModePerm)
	_ = os.Mkdir(serviceName+"/pb", os.ModePerm)
	_ = os.Mkdir(serviceName+"/pb/example", os.ModePerm)

	_ = util.CopyFolder("pb/example/", serviceName+"/pb/example/")

	_ = os.Mkdir(serviceName+"/client", os.ModePerm)
	CreateScaffoldScript()
	fmt.Println(serviceName, "Scaffolding ...")
	RunScaffold(serviceName)
	g.client = NewRPCClientGenerator(protoCoreFile, serviceName, serviceURL)
	g.service = NewServiceGenerator(serviceName, serviceURL, protoCoreFile)
	fmt.Println(serviceName, "Generating client ...")
	g.client.Generate()
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(serviceName, "Generating service&test ...")
	g.service.Generate()
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(serviceName, "Created")

}

//CreateScaffoldScript to create bash script for scaffolding
func CreateScaffoldScript() {
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
func GetTemplates(serviceName string) {
	contents := `#!/usr/bin/env bash
	servicename=$1;
	cd $servicename;
	git init;
	git remote add origin git@gitlab.kumparan.com:yowez/skeleton-service.git;
	git remote -v;
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
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

//GenerateProto2Go :nodoc:
func GenerateProto2Go(path string) {
	contents := `#!/usr/bin/env bash
	path=$1;
	path="${1}/*.proto";
	protoc --go_out=plugins=grpc:. $path;
	ls pb/example/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}';
`
	bt := []byte(contents)
	_ = ioutil.WriteFile(proto2go, bt, 0644)
	defer func() {
		_ = os.Remove(proto2go)
	}()
	pathArr := strings.Split(path,"/")
	pathArr = pathArr[:len(pathArr)-1]
	path = strings.Join(pathArr,"/")
	_= exec.Command(bash, proto2go, path )
}

//RunScaffold to run scaffold script
func RunScaffold(serviceName string) {
	exec.Command(bash, scaffold, serviceName)
	_ = os.Remove(scaffold)
}

func splitBetweenTwoChar(str, before, after string) string {
	a := strings.SplitAfterN(str, before, 2)
	b := strings.SplitAfterN(a[len(a)-1], after, 2)
	if 1 == len(b) {
		return b[0]
	}
	return b[0][0 : len(b[0])-len(after)]
}

func createSimpleNameFromProtoPath(str string) string {
	n := len(strings.Split(str, "/"))
	return strings.Split(str, "/")[n-2]
}
