package console

import (
	"github.com/kumparan/fer/gocek"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gocekCmd)
}

var gocekCmd = &cobra.Command{
	Use:   "gocek [target]",
	Short: "gocek check a module update info",
	Long:  `default check current working directory`,
	Run:   gocekRunner,
}

func gocekRunner(cmd *cobra.Command, args []string) {
	checker := gocek.ModuleChecker{}
	checker.CheckCWD()
}
