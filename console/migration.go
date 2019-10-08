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
