package console

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kumparan/fer/installer"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update fer",
	Long:  `update fer to latest version`,
	Run:   updateVersion,
}

func updateVersion(cmd *cobra.Command, args []string) {
	updateCommand := exec.Command("go", "get", "github.com/kumparan/fer@latest")
	updateCommand.Env = append(os.Environ(),
		"GO111MODULE=on",
	)
	err := updateCommand.Run()
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println("Failed updating fer")
		fmt.Println(err)
		os.Exit(1)
	}

	installer.ProgressBar(100)

	PrintSuccess("Success updating fer!")
}
