package console

import (
	"github.com/spf13/cobra"
)

// generatorCmd represents the generator command
var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "a generator",
	Long: `generate can generate project or db migration file or repository with model,
example 'fer generate [command]'`,
}

func init() {
	generatorCmd.AddCommand(projectCmd)
	generatorCmd.AddCommand(migrationCmd)
	generatorCmd.AddCommand(repositoryCmd)
}
