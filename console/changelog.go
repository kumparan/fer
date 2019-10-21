package console

import (
	"fmt"
	"log"

	"github.com/kumparan/fer/generator"
	"github.com/spf13/cobra"
)

func init() {
	generateInitChglogCmd.Flags().String("style", "none", "[optional] select the style according to your repository")
	initCmd.AddCommand(generateInitChglogCmd)
	createCmd.AddCommand(createChglogCmd)
}

// initCmd represents the generate command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "a initialization",
	Long: `initialization configuration, example 'fer init [command]'`,
}

var generateInitChglogCmd = &cobra.Command{
	Use:   "chglog [url]",
	Short: "Generate the changelog configuration file",
	Long:  "This subcommand to generate changelog ",
	Args:  cobra.ExactArgs(1),
	Run:   initChglog,
}

func initChglog(cmd *cobra.Command, args []string) {
	style, err := cmd.Flags().GetString("style")
	if err != nil {
		log.Fatal("fail retrieve --style flag")
	}
	if args[0] == "" {
		fmt.Println("please input repository URL 'fer generate chglog <style> <url>' ")
	}
	url := args[0]
	generator.InitChangelog(style, url)
}


// initCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "a create",
	Long: `created generate file, example 'fer create [command]'`,
}

var createChglogCmd = &cobra.Command{
	Use:   "chglog [version]",
	Short: "Created CHANGELOG.MD file",
	Long:  "This sub-command to create CHANGELOG.MD file to your repository",
	Args:  cobra.ExactArgs(1),
	Run:   createChglogMDFile,
}

func createChglogMDFile(cmd *cobra.Command, args []string) {
	if args[0] != "" {
		version := args[0]
		generator.CreateChangelog(version)
	}
<<<<<<< HEAD
	fmt.Println("error, version is not set")
=======
	fmt.Println("error version is not set")
>>>>>>> 0d8984088d14fd970665ec1acb7e3c22c3fb341b
}

