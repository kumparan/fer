package generator

import (
	"bufio"
	"bytes"
	"io/ioutil"

	"os"
	"strings"

	jen "github.com/dave/jennifer/jen" //Code generator
)

var (
	packageName         = "service"
	serviceURLProtoPath string
)

//GenerateServiceAndTest generate service&test from pb.go
func GenerateServiceAndTest(servicename string, prefixPackage string, pbCorePath string) {
	f, err := os.Open(pbCorePath)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(f)
	flagPrint := false
	serviceFlag := ""

	var service = make(map[string]string)
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
			service[strings.ToLower(serviceFlag)] += text + "\n"
		}
		if strings.Contains(scan.Text(), "Client interface") {
			flagPrint = true
		}

		if strings.Contains(scan.Text(), "}") {
			flagPrint = false
		}
	}

	for key, value := range service {
		f := jen.NewFile(packageName)
		fTest := jen.NewFile(packageName)
		serviceURLProtoPath = prefixPackage + "/pb/" + getSimpleNameFromProtoPath(pbCorePath)

		f.ImportAlias(serviceURLProtoPath, "pb")

		functions := strings.Split(value, "\n")

		for _, fun := range functions {
			if len(fun) < 5 {
				continue
			}
			if fun != "" {
				f.Comment(getFunctionName(fun) + " :nodoc:")
				f.Func().Params(
					jen.Id("s").Id("Service"),
				).Id(getFunctionName(fun)).Params(getFunctionArgs(fun)...).
					Parens(jen.List(jen.Id("res").Op("*").Qual(prefixPackage+"/pb/"+getSimpleNameFromProtoPath(pbCorePath), getFunctionReturns(fun)[1]),
						jen.Id("err").Error())).
					Block(jen.Return())

				//testfunction
				fTest.Func().Id(getFunctionTestName(fun)).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block()
			}
		}
		buf := &bytes.Buffer{}
		bufTest := &bytes.Buffer{}
		_ = f.Render(buf)
		_ = fTest.Render(bufTest)

		_ = ioutil.WriteFile(servicename+"/service/"+key+"_impl.go", buf.Bytes(), 0644)
		_ = ioutil.WriteFile(servicename+"/service/"+key+"_impl_test.go", bufTest.Bytes(), 0644)

	}
}

func getFunctionName(in string) (funcName string) {
	funcName = strings.Split(in, "(")[0]
	funcName = strings.ReplaceAll(funcName, ")", "")
	funcName = strings.TrimSpace(funcName)
	return
}
func getFunctionTestName(in string) (funcName string) {
	funcName = strings.Split(in, "(")[0]
	funcName = strings.ReplaceAll(funcName, ")", "")
	funcName = strings.TrimSpace(funcName)
	funcName = "Test" + funcName
	return
}

func getFunctionArgs(in string) (args []jen.Code) {
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
				argPath = serviceURLProtoPath
			}
			argType := strings.Split(argItem[1], ".")[1]
			args = append(args, jen.Code(jen.Id(argName).Op("*").Qual(argPath, argType)))
		}
	}

	return
}

func getFunctionReturns(in string) (rtns []string) {
	strlong := strings.Split(in, "(")[2]
	strlong = strings.ReplaceAll(strlong, ")", "")
	strs := strings.Split(strlong, ", ")
	for _, v := range strs {
		if strings.Contains(v, ".") {
			v = strings.ReplaceAll(v, "*", "")
			v = strings.ReplaceAll(v, ")", "")
			item := strings.Split(v, ".")
			rtns = append(rtns, item[0])
			rtns = append(rtns, item[1])
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

func getSimpleNameFromProtoPath(str string) string {
	n := len(strings.Split(str, "/"))
	return strings.Split(str, "/")[n-2]
}
