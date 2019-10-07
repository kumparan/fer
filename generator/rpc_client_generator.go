package generator

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	jen "github.com/dave/jennifer/jen" //Code generator
)

// Function :nodoc:
type Function struct {
	Name       string
	Parameters []string
	Returns    []string
}

type (
	//Client define client
	Client interface {
		Generate() error
	}
	client struct {
		protoPath         string
		serviceName       string
		serviceURL        string
		serviceURLProto   string
		serviceClientName string
	}
)

//NewRPCClientGenerator :nodoc:
func NewRPCClientGenerator(protoPath string, serviceName string, serviceRepo string) Client {
	return &client{
		serviceName:     serviceName,
		protoPath:       protoPath,
		serviceURL:      serviceRepo,
		serviceURLProto: serviceRepo + "/pb/" + createSimpleNameFromProtoPath(protoPath),
	}
}

//Generate ...
func (c client) Generate() error {
	f := jen.NewFile("client")
	serviceObj, err := c.ParseProtoToArray(c.serviceName, c.protoPath)
	if err != nil {
		return err
	}
	f.ImportAlias(c.serviceURLProto, "pb")
	for _, v := range serviceObj {
		var returns string
		var bodyReturn string
		parameters := "(" + strings.Join(v.Parameters, ",") + ")"

		for _, itemReturn := range v.Returns {
			returns = returns + "," + itemReturn
		}
		for _, itemReturn := range v.Returns {
			if strings.Contains(itemReturn, "*") {
				itemReturn = strings.Replace(itemReturn, "*", "&", -1)
				itemReturn += "{}"
			}
			if strings.Contains(itemReturn, "error") {
				itemReturn = strings.Replace(itemReturn, "error", "nil", -1)
			}
			bodyReturn = bodyReturn + "," + itemReturn

		}
		returns = returns[1:]
		returns = "( " + returns + " )"
		f.Func().Params(
			jen.Id("c").Op("*").Id("client"),
		).Id(v.Name).Params(c.CreateFunctionArgsClient(parameters)...).Id(returns).Block(
			c.CreateConn(),
			c.CreateErrChecker(),
			c.CreateCloseConn(),
			c.CreateNewClient(c.serviceClientName),
			jen.Return(c.CreateClientReturn(v.Name)),
		)
	}

	buf := &bytes.Buffer{}
	err = f.Render(buf)
	if err != nil {
		return err
	}
	splitPath := strings.Split(c.protoPath, "/")
	savePath := splitPath[2]
	savePath = strings.Replace(savePath, ".pb.go", ".go", -1)
	err = ioutil.WriteFile(c.serviceName+"/"+"client"+"/"+savePath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

//CreateFunctionArgsClient :nodoc:
func (c client) CreateFunctionArgsClient(in string) (args []jen.Code) {
	funcArgument := splitBetweenTwoChar(in, "(", ")")
	arguments := strings.Split(funcArgument, ", ")
	arguments = arguments[1 : len(arguments)-1]
	args = append(args, jen.Code(jen.Id("ctx").Qual("context", "Context")))
	for _, v := range arguments {
		if strings.Contains(v, ".") {
			argItem := strings.Split(v, " ")
			argName := argItem[0]
			argPath := strings.Split(argItem[1], ".")[0]
			if strings.Contains(argPath, "pb") {
				argPath = c.serviceURLProto
			}
			argType := strings.Split(argItem[1], ".")[1]
			args = append(args, jen.Code(jen.Id(argName).Op("*").Qual(argPath, argType)))
		}
	}
	args = append(args, jen.Code(jen.Id("opts").Op("...").Qual("google.golang.org/grpc", "CallOption")))

	return
}

// ParseProtoToArray ... :nodoc:
func (c client) ParseProtoToArray(serviceName string, path string) ([]Function, error) {
	interfaceName := "Client" + " interface"
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()
	scanner := bufio.NewScanner(f)

	isServiceClient := false
	services := []Function{}
	protoFunctions := []string{}
	for scanner.Scan() {
		var text string
		if strings.Contains(scanner.Text(), interfaceName) {
			isServiceClient = true
			c.serviceClientName = c.CreateServiceClientName(scanner.Text())
		}
		if isServiceClient {
			if strings.Contains(scanner.Text(), "//") {
				continue
			}
			if strings.Contains(scanner.Text(), "*") {
				text = scanner.Text()
				text = strings.Replace(text, "*", "*pb.", -1)
				protoFunctions = append(protoFunctions, text)
			}
		}
		if isServiceClient && strings.Contains(scanner.Text(), "}") {
			isServiceClient = false
		}
	}
	for _, v := range protoFunctions {
		params := c.SplitFunctionParameters(v)
		services = append(services, params)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return services, nil
}

//SplitFunctionParameters :nodoc:
func (c client) SplitFunctionParameters(function string) Function {
	splittedFunction := strings.Split(function, "(")
	for k := range splittedFunction {
		splittedFunction[k] = strings.Replace(splittedFunction[k], ")", "", -1)
	}
	params := strings.Split(splittedFunction[1], ",")
	returns := strings.Split(splittedFunction[2], ",")

	newFunction := Function{Name: splittedFunction[0], Parameters: params, Returns: returns}
	return newFunction
}

//CreateConn to generate conn code
func (c client) CreateConn() (s *jen.Statement) {
	s = jen.List(jen.Id("conn"), jen.Id("err")).Op(":=").Id("c").Dot("Conn").Dot("Get").Parens(jen.Id("ctx"))
	return
}

//CreateErrChecker to generate err checker
func (c client) CreateErrChecker() (s *jen.Statement) {
	s = jen.If(
		jen.Err().Op("!=").Nil(),
	).Block(
		jen.Return(jen.Nil(), jen.Err()),
	)
	return
}

//CreateCloseConn to generate close conn code
func (c client) CreateCloseConn() (s *jen.Statement) {
	s = jen.Defer().Func().Params().Block(
		jen.Id("_").Op("=").Id("conn").Dot("Close()")).Call()
	return
}

//CreateNewClient to generate new client grpc
func (c client) CreateNewClient(client string) (s *jen.Statement) {
	s = jen.Id("cli").Op(":=").Id("pb").Dot("New" + client).Parens(jen.Id("conn").Dot("ClientConn"))
	return
}

//CreateClientReturn to generate client return
func (c client) CreateClientReturn(funcName string) (s *jen.Statement) {
	s = jen.Id("cli").Dot(funcName).Parens(jen.List(jen.Id("ctx"), jen.Id("in"), jen.Id("opts...")))
	return
}

//CreateServiceClientName to get service client name
func (c client) CreateServiceClientName(text string) (client string) {
	client = text
	client = strings.ReplaceAll(client, "type", "")
	client = strings.ReplaceAll(client, "interface", "")
	client = strings.ReplaceAll(client, "{", "")
	client = strings.ReplaceAll(client, "\n", "")
	client = strings.ReplaceAll(client, " ", "")
	return
}
