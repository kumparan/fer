package console

import (
	"strings"

	"github.com/kumparan/fer/deployment"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deploymentCmd)
}

const (
	deployExample = `example "fer deploy dev-a"`
)

var deploymentCmd = &cobra.Command{
	Use:   "deploy [target]",
	Short: "service deployment",
	Long:  deployExample + "\nfor service deployment\n" + deployment.AvailableTargets,
	Args:  cobra.ExactArgs(1),
	Run:   deploymentService,
}

func deploymentService(cmd *cobra.Command, args []string) {
	dest := args[0]
	d := deployment.NewDeployment()
	d.Run(strings.ToLower(dest))
}
