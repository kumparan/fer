package generator

import (
	"bufio"
	"bytes"
	"io/ioutil"

	"os"
	"strings"

	jen "github.com/dave/jennifer/jen" //Code generator
)

const (
	packageName = "service"
)

type (
	//Service define service
	Service interface {
		Generate()
		CreateFunctionList() map[string]string
		CreateFunctionName(in string) string
		CreateTestFunctionName(in string) string
	}
	service struct {
		name                string
		url                 string
		protoPath           string
		serviceURLProtoPath string
	}
)

//NewServiceGenerator :nodoc:
func NewServiceGenerator(name string, url string, protoPath string) Service {
	return &service{
		name:      name,
		url:       url,
		protoPath: protoPath,
	}
}

//Generate generate service&test from proto
func (s service) Generate() {
	var functions = s.CreateFunctionList()
	for key, value := range functions {
		f := jen.NewFile(packageName)
		fTest := jen.NewFile(packageName)
		s.serviceURLProtoPath = s.url + "/pb/" + createSimpleNameFromProtoPath(s.protoPath)

		f.ImportAlias(s.serviceURLProtoPath, "pb")

		functions := strings.Split(value, "\n")

		for _, fun := range functions {
			if len(fun) < 5 {
				continue
			}
			if fun != "" {
				f.Comment(s.CreateFunctionName(fun) + " :nodoc:")
				f.Func().Params(
					jen.Id("s").Id("Service"),
				).Id(s.CreateFunctionName(fun)).Params(s.CreateFunctionArgs(fun)...).
					Parens(jen.List(jen.Id("res").Op("*").Qual(s.serviceURLProtoPath, s.CreateFunctionReturns(fun)[1]),
						jen.Id("err").Error())).
					Block(jen.Return())

				//testfunction
				fTest.Func().Id(s.CreateTestFunctionName(fun)).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block()
			}
		}
		buf := &bytes.Buffer{}
		bufTest := &bytes.Buffer{}
		_ = f.Render(buf)
		_ = fTest.Render(bufTest)

		_ = ioutil.WriteFile(s.name+"/service/"+key+"_impl.go", buf.Bytes(), 0644)
		_ = ioutil.WriteFile(s.name+"/service/"+key+"_impl_test.go", bufTest.Bytes(), 0644)

	}
}

//CreateFunctionList to create function list
func (s service) CreateFunctionList() (functions map[string]string) {
	functions = make(map[string]string)
	f, err := os.Open(s.protoPath)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(f)
	flagPrint := false
	serviceFlag := ""
	for scan.Scan() {
		if strings.Contains(scan.Text(), "TODO") {
			continue
		}
		if flagPrint {
			if strings.Contains(scan.Text(), "//") {
				serviceFlag = strings.Split(scan.Text(), " ")[1]
				continue
			}
			text := scan.Text()
			text = strings.ReplaceAll(text, "*", "*pb.")
			functions[strings.ToLower(serviceFlag)] += text + "\n"
		}
		if strings.Contains(scan.Text(), "Client interface") {
			flagPrint = true
		}

		if strings.Contains(scan.Text(), "}") {
			flagPrint = false
		}
	}
	return
}

//CreateFunctionName to create function name
func (s service) CreateFunctionName(in string) (funcName string) {
	funcName = strings.Split(in, "(")[0]
	funcName = strings.ReplaceAll(funcName, ")", "")
	funcName = strings.TrimSpace(funcName)
	return
}

//CreateTestFunctionname to create test function name
func (s service) CreateTestFunctionName(in string) (funcName string) {
	funcName = strings.Split(in, "(")[0]
	funcName = strings.ReplaceAll(funcName, ")", "")
	funcName = strings.TrimSpace(funcName)
	funcName = "Test" + funcName
	return
}

//CreateFunctionArgs to create function args
func (s service) CreateFunctionArgs(in string) (args []jen.Code) {
	strlong := splitBetweenTwoChar(in, "(", ")")
	strs := strings.Split(strlong, ", ")
	strs = strs[1 : len(strs)-1]

	args = append(args, jen.Code(jen.Id("ctx").Qual("context", "Context")))
	for _, v := range strs {
		if strings.Contains(v, ".") {
			argItem := strings.Split(v, " ")
			argName := argItem[0]
			argPath := strings.Split(argItem[1], ".")[0]
			if strings.Contains(argPath, "pb") {
				argPath = s.serviceURLProtoPath
			}
			argType := strings.Split(argItem[1], ".")[1]
			args = append(args, jen.Code(jen.Id(argName).Op("*").Qual(argPath, argType)))
		}
	}
	return
}

//CreateFunctionReturns to create function returns
func (s service) CreateFunctionReturns(in string) (r []string) {
	strlong := strings.Split(in, "(")[2]
	strlong = strings.ReplaceAll(strlong, ")", "")
	strs := strings.Split(strlong, ", ")
	for _, v := range strs {
		if strings.Contains(v, ".") {
			v = strings.ReplaceAll(v, "*", "")
			v = strings.ReplaceAll(v, ")", "")
			item := strings.Split(v, ".")
			r = append(r, item[0])
			r = append(r, item[1])
		}
	}
	return
}
