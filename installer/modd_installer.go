package installer

import (
	"fmt"
	"os"
	"os/exec"
)

// InstallModd :nodoc:
func InstallModd() {
	installModdCommand := exec.Command("go", "get", "github.com/cortesi/modd/cmd/modd@latest")
	installModdCommand.Env = append(os.Environ(),
		"GO111MODULE=on",
	)
	fmt.Println("installing modd")
	err := installModdCommand.Run()
	if err != nil {
		ProgressBar(1)
		fmt.Println("Failed installing mod")
		fmt.Println(err)
		os.Exit(1)
	}

	ProgressBar(100)

	fmt.Println("Success install modd!")
}
