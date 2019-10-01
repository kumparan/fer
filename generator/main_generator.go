package generator

import (
	"bufio"
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

//Generate to generate service
func Generate(servicename string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter proto path (Ex:\"pb/content/content.proto\"): ")
	text2, _ := reader.ReadString('\n')
	protoPath := strings.ReplaceAll(text2, "\n", "")
	protoPath = strings.ReplaceAll(protoPath, "proto", "pb.go")

	fmt.Println(">>>" + servicename + "<<<")
	fmt.Println("RPC Path : ", protoPath)

	protoCoreFile := protoPath
	serviceURL := "gitlab.kumparan.com/yowez/" + servicename

	fmt.Println("Creating ", servicename)

	_ = os.RemoveAll(servicename)
	_ = os.Mkdir(servicename, os.ModePerm)
	GetTemplates(servicename)
	_ = os.Mkdir(servicename+"/service", os.ModePerm)
	_ = os.Mkdir(servicename+"/pb", os.ModePerm)
	_ = os.Mkdir(servicename+"/pb/content", os.ModePerm)

	_ = util.CopyFolder("pb/content/", servicename+"/pb/content/")

	_ = os.Mkdir(servicename+"/client", os.ModePerm)
	GenerateProto2Go()
	CreateScaffoldScript()
	fmt.Println(servicename, "Scaffolding ...")
	RunScaffold(servicename)
	fmt.Println(servicename, "Generating client ...")
	GenerateRPCClient(protoCoreFile, servicename, serviceURL)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(servicename, "Generating service&test ...")
	GenerateServiceAndTest(servicename, serviceURL, protoCoreFile)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(servicename, "Created")

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
echo "finish scaffolding";

echo "running test";
make mockgen;
make test;
echo "finish test";

git init;
git remote add origin "git@gitlab.kumparan.com:yowez/$servicename.git";
echo "git initialized";
echo "Oke";
echo "finish";
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
	err := ioutil.WriteFile("gettemplate.sh", bt, 0644)
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("bash", "gettemplate.sh", serviceName)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	_ = os.Remove("gettemplate.sh")
}

//GenerateProto2Go :nodoc:
func GenerateProto2Go() {
	contents := `#!/usr/bin/env bash
	protoc --go_out=plugins=grpc:. pb/*.proto
	ls pb/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
`

	bt := []byte(contents)
	_ = ioutil.WriteFile("proto2go.sh", bt, 0644)
	exec.Command("bash", "proto2go.sh")
	_ = os.Remove("proto2go.sh")
}

//RunScaffold to run scaffold script
func RunScaffold(serviceName string) {
	cmd := exec.Command("bash", "scaffold.sh", serviceName)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	_ = os.Remove("scaffold.sh")
}
