package console

import (
	"log"
	"os"
	"path"

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
	rootDir := path.Join(GetConfigDir(), "gocek")
	_, err := os.Stat(rootDir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(rootDir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	checker := gocek.ModuleChecker{
		RootDir: rootDir,
	}
	// TODO: get dirs from some config files
	checker.Checks([]string{
		"/Users/miun/work/yowez/fer",
		"/Users/miun/work/yowez/document-service",
	})
}

func gocekCWD(cmd *cobra.Command, args []string) {
	checker := gocek.ModuleChecker{}
	checker.CheckCWD()
}
