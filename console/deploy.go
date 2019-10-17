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
	Use:   "deploy [target]",
	Short: "for service deployment, available target : beta,dev-a,dev-b,dev-c,dev-d,dev-e,preview,stable",
	Long:  `for service deployment, available target : beta,dev-a,dev-b,dev-c,dev-d,dev-e,preview,stable`,
	Args:  cobra.ExactArgs(1),

	Run: deployService,
}

func deployService(cmd *cobra.Command, args []string) {
	dest := args[0]
	deploy := dep.NewDeploy()
	deploy.Run(strings.ToLower(dest))
}
