package generator

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dave/jennifer/jen" //Code project
)

const (
	packageName = "service"
)

type (
	//Service define service
	Service interface {
		Generate() error
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
func (s service) Generate() error {
	var functions, err = s.CreateFunctionList()
	if err != nil {
		return err
	}
	for key, value := range functions {
		f := jen.NewFile(packageName)
		fTest := jen.NewFile(packageName)
		s.serviceURLProtoPath = path.Join(s.url, "pb", createSimpleNameFromProtoPath(s.protoPath))

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
				).Id(s.CreateFunctionName(fun)).Params(s.CreateFunctionParameters(fun)...).
					Parens(jen.List(jen.Id("res").Op("*").Qual(s.serviceURLProtoPath, s.CreateFunctionReturns(fun)[1]),
						jen.Id("err").Error())).
					Block(jen.Return())

				//testfunction
				fTest.Func().Id(s.CreateTestFunctionName(fun)).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block()
			}
		}

		buf := &bytes.Buffer{}
		bufTest := &bytes.Buffer{}
		err := f.Render(buf)
		if err != nil {
			return err
		}
		err = fTest.Render(bufTest)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path.Join(s.name, "service", key+"_impl.go"), buf.Bytes(), 0644)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(path.Join(s.name, "service", key+"_impl_test.go"), bufTest.Bytes(), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

//CreateFunctionList to create function list
func (s service) CreateFunctionList() (functions map[string]string, err error) {
	functions = make(map[string]string)
	f, err := os.Open(s.protoPath)
	if err != nil {
		return nil, err
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

//CreateFunctionParameters to create function parameters
func (s service) CreateFunctionParameters(in string) (parameters []jen.Code) {
	param := splitBetweenTwoChar(in, "(", ")") //get all parameter
	params := strings.Split(param, ", ")       //split each parameter
	params = params[1 : len(params)-1]         //remove grpc option

	parameters = append(parameters, jen.Code(jen.Id("ctx").Qual("context", "Context")))
	for _, v := range params {
		if strings.Contains(v, ".") {
			paramItem := strings.Split(v, " ")
			paramName := paramItem[0]
			paramPath := strings.Split(paramItem[1], ".")[0]
			if strings.Contains(paramPath, "pb") {
				paramPath = s.serviceURLProtoPath
			}
			paramType := strings.Split(paramItem[1], ".")[1]
			parameters = append(parameters, jen.Code(jen.Id(paramName).Op("*").Qual(paramPath, paramType)))
		}
	}
	return
}

//CreateFunctionReturns to create function returns
func (s service) CreateFunctionReturns(in string) (returns []string) {
	ret := strings.Split(in, "(")[2] //get all return
	ret = strings.ReplaceAll(ret, ")", "")
	retSlice := strings.Split(ret, ", ")
	for _, v := range retSlice {
		if strings.Contains(v, ".") {
			v = strings.ReplaceAll(v, "*", "")
			v = strings.ReplaceAll(v, ")", "")
			item := strings.Split(v, ".")
			returns = append(returns, item[0])
			returns = append(returns, item[1])
		}
	}
	return
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
