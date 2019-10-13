package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/kumparan/fer/installer"
)

// InitChangelog :nodoc:
func InitChangelog(style, repositoryURL string) {
	fmt.Println("Generate changelog configuration")
	getWorkingDirectory, err := exec.Command("pwd").Output()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("fail generate chglog")
		os.Exit(1)
	}
	workingDirectory := string(getWorkingDirectory)
	cmd := exec.Command("cd", workingDirectory)
	err = cmd.Run()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("fail generate chglog")
		os.Exit(1)
	}

	err = os.MkdirAll(".chglog", os.ModePerm)
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("fail generate chglog")
		os.Exit(1)
	}

	// message := []byte("Hello, Gophers!")
	err = ioutil.WriteFile(".chglog/config.yml", []byte(installer.CreateChangelogyml(style, repositoryURL)), 0644)
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("fail generate chglog")
		os.Exit(1)
	}

	err = ioutil.WriteFile(".chglog/CHANGELOG.tpl.md", []byte(installer.CreateChangelogMD()), 0644)
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("fail generate chglog")
		os.Exit(1)
	}
	installer.ProgressBar(100)
	fmt.Println("Success generate chglog")
	os.Exit(0)
}
