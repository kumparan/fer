package cmd

import (
	"fmt"

	"github.com/kumparan/fer/generator"
	"github.com/spf13/cobra"
)

// projectCmd represents the init command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "project",
	Long:  `example 'fer project --name content-service',and then input the service proto path`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name != "" {
			generator.Generate(name)
		} else {
			fmt.Println("please add --name 'example-service' for service name")
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().String("name", "", "name for new microservice")
}
