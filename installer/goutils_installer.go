package installer

import (
	"fmt"
	"os/exec"
)

// InstallGoUtils :nodoc:
func InstallGoUtils(installername, URL string) string {
	cmd := exec.Command("go", "get", "-u", URL)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return "fail installed " + installername
	}
	message := CheckedInstallerPath(installername)
	return message

}
