package console

import (
	"strings"

	dep "github.com/kumparan/fer/deploy"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploying service",
	Long:  `deploy service`,
	Args:  cobra.ExactArgs(1),
	Run:   deployService,
}

func deployService(cmd *cobra.Command, args []string) {
	dest := args[0]
	deploy := dep.NewDeploy()
	deploy.Run(strings.ToLower(dest))
}
