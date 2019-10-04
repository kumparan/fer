package installer

import (
	"fmt"
	"os"
	"os/exec"
)

type (
	// WatchmedoInstaller :nodoc:
	WatchmedoInstaller interface {
		InstallWatchmedo()
	}
)

// InstallWatchmedo :nodoc:
func InstallWatchmedo() string {
	pipCmd := "pip3"
	cmdGetPip3Location := exec.Command("which", pipCmd)
	err := cmdGetPip3Location.Run()
	if err != nil {
		pipCmd = "pip"
		cmdGetPipLocation := exec.Command("which", pipCmd)
		err = cmdGetPipLocation.Run()
		if err != nil {
			fmt.Println("you must install python-pip first")
			os.Exit(1)
		}

	}
	cmdInstallWachmedoByPip := exec.Command(pipCmd, "install", "watczhdog")
	err = cmdInstallWachmedoByPip.Run()
	if err != nil {
		fmt.Println(err)
		return "fail installed watchmedo"
	}
	message := CheckedInstallerPath("watchmedo")
	return message
}
