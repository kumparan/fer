package console

import (
	"fmt"
	"log"

	"github.com/kumparan/fer/generator"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "project",
	Long: `fer project is a microservice generator, you can generate microservice from proto.
example 'fer project example-service --proto pb/example/example.proto' 
		`,
	Args: cobra.ExactArgs(1),
	Run:  projectGenerator,
}

func projectGenerator(cmd *cobra.Command, args []string) {
	g := generator.NewGenerator()
	name := args[0]
	proto, err := cmd.Flags().GetString("proto")
	if err != nil {
		log.Fatal("fail retreive --proto flag")
	}
	if name != "" {
		g.Run(name, proto)
	} else {
		fmt.Println("you must add name in 'fer project <name>' for service name, if you had proto include using this --proto 'pb/example/example.proto' for proto path")
	}
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().String("proto", "", "[optional] proto path to generate service")
}
