package installer

import (
	"fmt"
	"os/exec"
)

// InstallGoUtils :nodoc:
func InstallGoUtils(installername, URL string) {
	fmt.Println("Installing ", installername)
	cmd := exec.Command("go", "get", "-u", URL)
	err := cmd.Run()
	if err != nil {
		ProgressBar(1)
		fmt.Println(err)
		fmt.Println("fail installed " + installername)
	}
	ProgressBar(100)
	message := CheckedInstallerPath(installername)
	fmt.Println(message)
}
