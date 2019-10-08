package console

import (
	"fmt"
	"log"

	"github.com/kumparan/fer/generator"
	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration [name]",
	Short: "migration",
	Long: `Create DB Migration file,
example 'fer migration create_promoted_link'`,
	Args: cobra.ExactArgs(1),
	Run:  migrationGenerator,
}

func migrationGenerator(cmd *cobra.Command, args []string) {
	if args[0] != "" {
		m := generator.NewMigrationGenerator()
		err := m.Generate(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("please input name 'fer migration <name>' ")
	}
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
