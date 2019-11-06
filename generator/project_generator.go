package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/kumparan/fer/util"

	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
)

const (
	scaffold = "scaffold.sh"
	proto2go = "proto2go.sh"
	template = "template.sh"
	bash     = "bash"
)

type (
	//Project define project
	Project interface {
		Run(serviceName, protoPath string)
	}
	project struct {
		service     Service
		client      Client
		protoFile   string
		serviceName string
	}
)

//NewProjectGenerator :nodoc:
func NewProjectGenerator() Project {
	return &project{}
}

//Generate to generate service
func (g project) Run(serviceName string, protoPath string) {
	g.serviceName = serviceName
	fmt.Println(">>> " + serviceName + " <<<")
	fmt.Println("Creating ", serviceName)
	fmt.Println(serviceName, "Scaffolding ...")
	err := os.Mkdir(serviceName, os.ModePerm)
	if err != nil {
		log.Fatal("folder " + serviceName + " already exist")
	}
	err = os.Mkdir(path.Join(serviceName, "client"), os.ModePerm)
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
		err = util.CopyFolder(g.getProtoFolder(protoPath)+"/", path.Join(serviceName, g.getProtoFolder(protoPath)))
		if err != nil {
			g.rollbackWhenError("fail to create dir " + path.Join(serviceName, g.getProtoFolder(protoPath)))
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
	err = g.changeServiceNameOnMakefile(serviceName)
	if err != nil {
		g.rollbackWhenError("fail generate makefile " + err.Error())
	}
	err = g.removeWorker(serviceName)
	if err != nil {
		g.rollbackWhenError("fail to remove worker " + err.Error())
	}
	err = g.removeHello(serviceName)
	if err != nil {
		g.rollbackWhenError("fail to remove example files " + err.Error())
	}
	err = g.generateClient(serviceName)
	if err != nil {
		g.rollbackWhenError("fail to remove worker " + err.Error())
	}
	fmt.Println(serviceName, "Created")
}

//CreateScaffoldScript to create bash script for scaffolding
func (g project) CreateScaffoldScript() {
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
func (g project) GetTemplates(serviceName string) {
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
func (g project) GenerateProto2Go(path string) {
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
func (g project) RunScaffold(serviceName string) {
	cmd := exec.Command(bash, scaffold, serviceName)
	defer func() {
		_ = os.Remove(scaffold)
	}()
	err := g.runCmd(cmd)
	if err != nil {
		g.rollbackWhenError("fail on scaffolding")
	}
}

func (g project) runCmd(cmd *exec.Cmd) error {
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (g project) getProtoFolder(path string) string {
	pathArr := strings.Split(path, "/")
	pathArr = pathArr[:len(pathArr)-1]
	path = strings.Join(pathArr, "/")
	return path
}

func (g project) rollbackWhenError(message string) {
	_ = os.RemoveAll(g.serviceName)
	_ = os.Remove(scaffold)
	_ = os.Remove(template)
	_ = os.Remove(proto2go)
	log.Fatal(message)
}

func (g project) changeServiceNameOnMakefile(serviceName string) error {
	makefilePath := filepath.Join(serviceName, "Makefile")
	data, err := ioutil.ReadFile(makefilePath)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(makefilePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	newMakefile := strings.Replace(string(data), "skeleton", serviceName, -1)
	_, err = file.WriteString(newMakefile)
	if err != nil {
		return err
	}
	return nil
}

func (g project) removeHello(serviceName string) (err error) {
	err = os.Remove(path.Join(serviceName, "service", "hello_service_impl.go"))
	err = os.Remove(path.Join(serviceName, "service", "hello_service_impl_test.go"))
	err = os.Remove(path.Join(serviceName, "repository", "hello_repository_test.go"))
	err = os.Remove(path.Join(serviceName, "repository", "hello_repository.go"))
	err = os.Remove(path.Join(serviceName, "repository", "model/greeting.go"))
	err = os.RemoveAll(path.Join(serviceName, "pb", "skeleton"))
	return
}

func (g project) removeWorker(serviceName string) error {
	os.RemoveAll(path.Join(serviceName, "worker"))
	os.RemoveAll(path.Join(serviceName, "event"))
	os.Remove(path.Join(serviceName, "service", "service.go"))
	os.Remove(path.Join(serviceName, "console", "worker.go"))
	contents := `package service

import (
	"github.com/kumparan/go-lib/redcachekeeper"
	"github.com/kumparan/kumnats"
)

// Service :nodoc:
type Service struct {
	nats        kumnats.NATS
	cacheKeeper redcachekeeper.Keeper
}

// RegisterNATS :nodoc:
func (s *Service) RegisterNATS(n kumnats.NATS) {
	s.nats = n
}

// GetNATS :nodoc:
func (s *Service) GetNATS() kumnats.NATS {
	return s.nats
}

// NewSkeletonService :nodoc:
func NewSkeletonService() *Service {
	return new(Service)
}


// RegisterCacheKeeper :nodoc:
func (s *Service) RegisterCacheKeeper(k redcachekeeper.Keeper) {
	s.cacheKeeper = k
}
	`
	contents = strings.Replace(contents, "Skeleton", serviceName, -1)
	bt := []byte(contents)
	err := ioutil.WriteFile(path.Join(serviceName, "service", "service.go"), bt, 0644)
	if err != nil {
		g.rollbackWhenError("fail to write service")
	}
	return nil
}

func (g project) generateClient(serviceName string) error {
	os.Remove(path.Join(serviceName, "client", "client.go"))
	contents := `package client
	
	import (
		"context"
		"time"
	
		grpcpool "github.com/processout/grpc-go-pool"
		log "github.com/sirupsen/logrus"
		"google.golang.org/grpc"
	)
	
	type client struct {
		Conn *grpcpool.Pool
	}
	
	//NewClient is a func to create Client
	func NewClient(target string, timeout time.Duration, idleConnPool, maxConnPool int) (pb.HelloServiceClient, error) {
		factory := newFactory(target, timeout)
	
		pool, err := grpcpool.New(factory, idleConnPool, maxConnPool, time.Second)
		if err != nil {
			log.Errorf("Error : %v", err)
			return nil, err
		}
	
		return &client{
			Conn: pool,
		}, nil
	}
	
	func newFactory(target string, timeout time.Duration) grpcpool.Factory {
		return func() (*grpc.ClientConn, error) {
			conn, err := grpc.Dial(target, grpc.WithInsecure(), withClientUnaryInterceptor(timeout))
			if err != nil {
				log.Errorf("Error : %v", err)
				return nil, err
			}
	
			return conn, err
		}
	}
	
	func withClientUnaryInterceptor(timeout time.Duration) grpc.DialOption {
		return grpc.WithUnaryInterceptor(func(
			ctx context.Context,
			method string,
			req interface{},
			reply interface{},
			cc *grpc.ClientConn,
			invoker grpc.UnaryInvoker,
			opts ...grpc.CallOption,
		) error {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			err := invoker(ctx, method, req, reply, cc, opts...)
			return err
		})
	}
	`
	serviceNameOnly := strings.Replace(serviceName, "-service", "", -1)
	contents = strings.Replace(contents, "Skeleton", strcase.ToCamel(serviceNameOnly), -1)
	bt := []byte(contents)
	err := ioutil.WriteFile(path.Join(serviceName, "client", "client.go"), bt, 0644)
	if err != nil {
		g.rollbackWhenError("fail to write client")
	}
	return nil
}
