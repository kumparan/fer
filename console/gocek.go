package console

import (
	"github.com/kumparan/fer/config"
	"github.com/kumparan/fer/gocek"
	"github.com/spf13/cobra"
)

func init() {
	gocekCmd.AddCommand(gocekAllCmd)
	rootCmd.AddCommand(gocekCmd)
}

var gocekCmd = &cobra.Command{
	Use:   "gocek",
	Short: "gocek check a module update info",
	Long:  `default check current working directory`,
	Run:   gocekCWD,
}

var gocekAllCmd = &cobra.Command{
	Use:   "all",
	Short: "check all module update info",
	Run:   gocekAll,
}

func gocekAll(cmd *cobra.Command, args []string) {
	cfg := config.GetFerConfig()
	checker := gocek.ModuleChecker{
		RootDir:      cfg.Gocek.SaveOutputDir,
		ServicesDirs: cfg.Gocek.ProjectDirs,
	}

	checker.Checks()
}

func gocekCWD(cmd *cobra.Command, args []string) {
	checker := gocek.ModuleChecker{}
	checker.CheckCWD()
}
