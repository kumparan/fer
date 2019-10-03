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
	Long: `fer project is a microservice generator, you can generate microservice from proto.
example 'fer project --name example-service --proto pb/example/example.proto' 
		`,
	Run: projectGenerator,
}

func projectGenerator(cmd *cobra.Command, args []string) {
	g := generator.NewGenerator()
	name, _ := cmd.Flags().GetString("name")
	proto, _ := cmd.Flags().GetString("proto")
	if name != "" {
		g.Run(name, proto)
	} else {
		fmt.Println("you must add --name 'example-service' for service name, if you had proto include using this --proto 'pb/example/example.proto' for proto path")
	}
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().String("name", "", "name for new microservice")
	projectCmd.Flags().String("proto", "", "proto path to generate service")
}
