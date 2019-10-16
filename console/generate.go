package console

import (
	"fmt"
	"log"

	"github.com/kumparan/fer/generator"
	"github.com/spf13/cobra"
)

func init() {
	projectCmd.Flags().String("proto", "", "[optional] proto path to generate service")
	generateCmd.AddCommand(projectCmd)
	generateCmd.AddCommand(migrationCmd)
	generateCmd.AddCommand(repositoryCmd)
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "a generator",
	Long: `generate can generate project or db migration file or repository with model,
example 'fer generate [command]'`,
}

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "project generator",
	Long: `project is a microservice generator, you can generate microservice from proto.
example 'fer generate project example-service --proto pb/example/example.proto' 
		`,
	Args: cobra.ExactArgs(1),
	Run:  generateProject,
}

func generateProject(cmd *cobra.Command, args []string) {
	g := generator.NewProjectGenerator()
	name := args[0]
	proto, err := cmd.Flags().GetString("proto")
	if err != nil {
		log.Fatal("fail retreive --proto flag")
	}
	if name != "" {
		g.Run(name, proto)
	} else {
		fmt.Println("you must add name in 'fer generate project <name>' for service name, if you had proto include using this --proto 'pb/example/example.proto' for proto path")
	}
}

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use:   "repository [name]",
	Short: "repository generator",
	Long: `Generate repository and model file,
example 'fer generate repository promoted_link'`,
	Args: cobra.ExactArgs(1),
	Run:  generateRepository,
}

func generateRepository(cmd *cobra.Command, args []string) {
	if args[0] != "" {
		r := generator.NewRepositoryGenerator()
		err := r.Generate(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("please input name 'fer generate repository <name>' ")
	}
}

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration [name]",
	Short: "migration generator",
	Long: `Create DB Migration file,
example 'fer generate migration create_promoted_link'`,
	Args: cobra.ExactArgs(1),
	Run:  generateMigration,
}

func generateMigration(cmd *cobra.Command, args []string) {
	if args[0] != "" {
		m := generator.NewMigrationGenerator()
		err := m.Generate(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("please input name 'fer generate migration <name>' ")
	}
}