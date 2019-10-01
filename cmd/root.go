package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-example",
	Short: "An example of cobra",
	Long: `This application shows how to create modern CLI
			applications in go using Cobra CLI library`,
}

// Execute :nodoc:
func Execute() {
	fmt.Printf(`
   ________________     __              __
   / ____/ ____/ __ \   / /_____  ____  / /
  / /_  / __/ / /_/ /  / __/ __ \/ __ \/ / 
 / __/ / /___/ _, _/  / /_/ /_/ / /_/ / /  
/_/   /_____/_/ |_|   \__/\____/\____/_/  

`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
