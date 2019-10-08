package generator

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
)

type (
	//Repository define repository generator
	Repository interface {
		Generate(name string) error
	}

	repository struct{}
)

//NewRepositoryGenerator :nodoc:
func NewRepositoryGenerator() Repository {
	return &repository{}
}

//GenerateRepository :nodoc:
func (r repository) Generate(name string) error {
	err := r.checkFolderExists()
	if err != nil {
		return err
	}
	err = r.generateModel(name)
	if err != nil {
		return err
	}
	fileName := strings.ToLower(name) + "_repository"
	structName := strings.ToLower(r.toCamelCase(name)) + "Repo"
	interfaceName := r.toCamelCase(name) + "Repository"
	repo := jen.NewFile("repository")

	repo.ImportAlias("github.com/jinzhu/gorm", "gorm")
	repo.ImportAlias("github.com/kumparan/cacher", "cacher")

	repo.Comment(interfaceName + " :nodoc:")
	repo.Type().Id(interfaceName).Interface(
		jen.Id("Create").Params(jen.Id("ctx").Qual("context", "Context")).Error(),
		jen.Id("FindByID").Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("id").String()).Error(),
		jen.Id("UpdateByID").Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("id").String()).Error(),
		jen.Id("DeleteByID").Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("id").String()).Error(),
		jen.Id("FindAll").Params(jen.Id("ctx").Qual("context", "Context")).Error(),
	)

	repo.Type().Id(structName).Struct(
		jen.Id("db").Op("*").Qual("github.com/jinzhu/gorm", "DB"),
		jen.Id("cacher").Qual("github.com/kumparan/cacher", "Keeper"),
	)

	repo.Comment("New" + r.toCamelCase(name) + "Repository" + " create new repository")
	repo.Func().Id("New" + r.toCamelCase(name) + "Repository").Params(jen.List(
		jen.Id("d").Op("*").Qual("github.com/jinzhu/gorm", "DB"),
		jen.Id("k").Qual("github.com/kumparan/cacher", "Keeper")),
	).Id(interfaceName).Block(
		jen.Return(jen.Op("&").Id(structName).Values(jen.Dict{
			jen.Id("db"):          jen.Id("d"),
			jen.Id("cacher"): jen.Id("k"),
		}),
		),
	)

	buf := &bytes.Buffer{}
	err = repo.Render(buf)
	if err != nil {
		return err
	}

	repositoryFileName := "repository/" + fileName + ".go"
	err = ioutil.WriteFile(repositoryFileName, buf.Bytes(), 0666)
	if err != nil {
		return err
	}
	fmt.Println(repositoryFileName + " created")

	return nil
}

func (r repository) generateModel(name string) error {
	m := jen.NewFile("model")
	modelName := r.toCamelCase(name)

	m.Comment(modelName + " represent " + strings.ReplaceAll(name, "_", " "))
	m.Type().Id(modelName).Struct(
		jen.Id("id").String(),
	)

	buf := &bytes.Buffer{}
	err := m.Render(buf)
	if err != nil {
		return err
	}

	modelFileName := "repository/model/" + strings.ToLower(name) + ".go"
	err = ioutil.WriteFile(modelFileName, buf.Bytes(), 0666)
	if err != nil {
		return err
	}
	fmt.Println(modelFileName + " created")
	return nil
}

func (r repository) checkFolderExists() error {
	_, err := os.Stat("repository/model/")
	if os.IsNotExist(err) { //check if repository/model/ folder is already exist
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("repository/model/ folder not found, want to  create (Y/N)? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return errors.New("fail when read user input")
		}
		ans := strings.Contains(strings.ToUpper(input), "Y")
		if !ans {
			return errors.New("cancel create repository & model")
		}
		err = r.createRepositoryFolder()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r repository) createRepositoryFolder() error {
	_, err := os.Stat("repository/")
	if os.IsNotExist(err) { //check if repository folder is already exist
		err := os.Mkdir("repository/", os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = os.Mkdir("repository/model/", os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) toCamelCase(str string) string {
	str = strings.ToLower(str)
	var regex = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
	return regex.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}
