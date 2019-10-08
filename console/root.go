package console

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fer",
	Short: "fer",
	Long: `
    ________________ 
   / ____/ ____/ __ \
  / /_  / __/ / /_/ /
 / __/ / /___/ _, _/ 
/_/   /_____/_/ |_|  
fer is not ferdian.
fer is development tool for backend engineer.
fer can help whatever you want.
`,
}

// Execute :nodoc:
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generatorCmd)
}
