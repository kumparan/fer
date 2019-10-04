package installer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/kumparan/fer/config"

	version "github.com/hashicorp/go-version"
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
	var regexVersion, _ = regexp.Compile(`(\d+\.\d+\.\d+)`)
	v1, _ := version.NewVersion(config.GoVersion)
	v2, _ := version.NewVersion(regexVersion.FindString(goLocalversion))
	if v2.LessThan(v1) {
		fmt.Printf("Go version must be %s or latest\n", config.GoVersion)
		os.Exit(1)
	}
}

// CheckedInstallerPath :nodoc:
func CheckedInstallerPath(installer string) string {
	cmdGetInstallerPath := exec.Command("which", installer)
	err := cmdGetInstallerPath.Run()
	if err != nil {
		return "fail installed " + installer
	}
	return "Success installed " + installer
}