package console

import (
	"fmt"
	"log"

	"github.com/kumparan/fer/generator"
	"github.com/spf13/cobra"
)

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use:   "repository [name]",
	Short: "repository",
	Long: `Generate repository and model file,
example 'fer repository promoted_link'`,
	Args: cobra.ExactArgs(1),
	Run:  repositoryGenerator,
}

func repositoryGenerator(cmd *cobra.Command, args []string) {
	if args[0] != "" {
		r := generator.NewRepositoryGenerator()
		err := r.Generate(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("please input name 'fer repository <name>' ")
	}
}

func init() {
	rootCmd.AddCommand(repositoryCmd)
}
