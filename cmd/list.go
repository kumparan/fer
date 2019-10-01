package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the task",
	Long:  "This subcommand start the test",
	Run:   list,
}

// Execute :nodoc:
func init() {
	RootCmd.AddCommand(listCmd)
}

func list(_ *cobra.Command, _ []string) {
	fmt.Println("list called")
}
