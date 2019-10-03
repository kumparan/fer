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
func InstallWatchmedo() {
	cmdGetPip3Location := exec.Command("which", "pip3")
	err := cmdGetPip3Location.Run()
	if err != nil {
		cmdGetPipLocation := exec.Command("which", "pip")
		err = cmdGetPipLocation.Run()
		if err != nil {
			fmt.Println("you must install python first")
			os.Exit(1)
		}
		cmdInstallWachmedoByPip := exec.Command("pip", "install", "watchdog")
		cmdInstallWachmedoByPip.Stdout = os.Stdout
		cmdInstallWachmedoByPip.Stderr = os.Stderr

		err = cmdInstallWachmedoByPip.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
	cmdInstallWachmedoByPip3 := exec.Command("pip3", "install", "watchdog")
	cmdInstallWachmedoByPip3.Stdout = os.Stdout
	cmdInstallWachmedoByPip3.Stderr = os.Stderr

	err = cmdInstallWachmedoByPip3.Run()
	if err != nil {
		fmt.Println(err)
	}
}
