package console

import (
	"fmt"

	"github.com/kumparan/fer/config"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  `print version of fer`,
	Run:   versionPrint,
}

func versionPrint(cmd *cobra.Command, args []string) {
	fmt.Println(config.Version)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
