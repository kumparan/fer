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
	cmd := exec.Command("git", "ls-remote", repositoryURL)
	err := cmd.Run()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("The project you were looking for could not be found.")
		fmt.Println(err)
		os.Exit(1)
	}

	getWorkingDirectory, err := exec.Command("pwd").Output()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}

	workingDirectory := string(getWorkingDirectory)
	cmd = exec.Command("cd", workingDirectory)
	err = cmd.Run()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := os.Stat(".chglog/config.yml"); !os.IsNotExist(err) {
		if err != nil {
			installer.ProgressBar(1)
			fmt.Println("failed to generate changelog")
			fmt.Println(err)
			os.Exit(1)
		}
		installer.ProgressBar(1)
		fmt.Println("chglog configuration already exist")
		os.Exit(1)
	}

	err = os.MkdirAll(".chglog", os.ModePerm)
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(".chglog/config.yml", []byte(installer.CreateChangelogYAML(style, repositoryURL)), 0644)
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(".chglog/CHANGELOG.tpl.md", []byte(installer.CreateChangelogMD()), 0644)
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}
	installer.ProgressBar(100)
	fmt.Println("changelog is succesfully created")
	os.Exit(0)
}

//CreateChangelog :nodoc:
func CreateChangelog(version string){
	fmt.Println("Creating CHANGELOG.md file")
	getWorkingDirectory, err := exec.Command("pwd").Output()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}

	workingDirectory := string(getWorkingDirectory)
	cmd := exec.Command("cd", workingDirectory)
	err = cmd.Run()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}

	createChgLogFile := exec.Command("git-chglog", "--next-tag", version, "-o", "CHANGELOG.md", "-p", "^v")
	err = createChgLogFile.Run()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("failed to generate changelog")
		fmt.Println(err)
		os.Exit(1)
	}
	installer.ProgressBar(100)
	fmt.Println("CHANGELOG.md is succesfully created")
	os.Exit(0)
}
