package repository

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/kumparan/fer/config"

	version "github.com/hashicorp/go-version"
)

type (
	// InstallRepository :nodoc:
	InstallRepository interface {
		CheckExistenceOfGolang()
		CheckGolangVersion()
		CheckExistenceOfChangelog()
		CheckChangelogVersion()
		InstalltheChangelog()
	}
)

// CheckExistenceOfGolang :nodoc:
func CheckExistenceOfGolang() {
	cmdGetGolangLocation := exec.Command("which", "go")
	err := cmdGetGolangLocation.Run()
	if err != nil {
		fmt.Println("You should install golang first")
		os.Exit(1)
	}
}

// CheckGolangVersion :nodoc:
func CheckGolangVersion() {
	cmdGetGolangVersion, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	var goLocalversion = string(cmdGetGolangVersion)
	var regexVersion, _ = regexp.Compile(`\d+(\.\d+){2,}`)
	v1, _ := version.NewVersion(config.GoVersion)
	v2, _ := version.NewVersion(regexVersion.FindString(goLocalversion))
	if v2.LessThan(v1) {
		fmt.Println("Go version must be 1.12.7 or latest")
		os.Exit(1)
	}
}

// CheckExistenceOfChangelog :nodoc:
func CheckExistenceOfChangelog() {
	cmdGetChglogLocation := exec.Command("which", "git-chglog")
	err := cmdGetChglogLocation.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// CheckChangelogVersion :nodoc:
func CheckChangelogVersion() {
	cmdGetChglogVersion, err := exec.Command("git-chglog", "--version").Output()
	if err != nil {
		log.Fatal(err)
	}
	var goLocalversion = string(cmdGetChglogVersion)
	fmt.Println(goLocalversion)
}

// InstalltheChangelog :nodoc:
func InstalltheChangelog() {
	cmd := exec.Command("go", "get", "-u", "github.com/git-chglog/git-chglog/cmd/git-chglog")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
