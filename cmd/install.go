package cmd

import (
	"fmt"
	"strings"

	"github.com/kumparan/fer/config"
	"github.com/kumparan/fer/installer"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install all dependencies",
	Long:  "install all dependencies for contributing to backend projects",
	Run:   installAllCmd,
}

func installAllCmd(_ *cobra.Command, _ []string) {
	installAll()
}

func installAll() {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	var messages = []string{}

	message := installer.InstallGoUtils("protoc-gen-go", config.ProtobufInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("mockgen", config.MockgenInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("richgo", config.RichgoInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("golint", config.GolintInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("git-chglog", config.ChangeLogInstallerURL)
	messages = append(messages, message)

	message = installer.CheckedInstallerPath("make")
	messages = append(messages, message)

	message = installer.InstallWatchmedo()
	messages = append(messages, message)

	message = installer.ProtobufInstaller()
	messages = append(messages, message)
	fmt.Printf("%s", strings.Join(messages, "/n"))

}

func init() {
	installCmd.AddCommand(goUtilsCmd)
	installCmd.AddCommand(watchmedoCmd)
	rootCmd.AddCommand(installCmd)
}

var goUtilsCmd = &cobra.Command{
	Use:   "goutils",
	Short: "fer install goutils",
	Long:  "This subcommand to install git go utils ",
	Run:   installGoUtilsCmd,
}

func installGoUtilsCmd(_ *cobra.Command, _ []string) {
	installGoUtils()
}

func installGoUtils() {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	var messages = []string{}
	message := installer.InstallGoUtils("protoc-gen-go", config.ProtobufInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("mockgen", config.MockgenInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("richgo", config.RichgoInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("golint", config.GolintInstallerURL)
	messages = append(messages, message)

	message = installer.InstallGoUtils("git-chglog", config.ChangeLogInstallerURL)
	messages = append(messages, message)

	message = installer.CheckedInstallerPath("make")
	messages = append(messages, message)

	fmt.Printf("%s", strings.Join(messages, "/n"))
}

var watchmedoCmd = &cobra.Command{
	Use:   "watchmedo",
	Short: "fer install watchmedo",
	Long:  "This subcommand to install watchmedo",
	Run:   installWatchmedoCmd,
}

func installWatchmedoCmd(cmd *cobra.Command, args []string) {
	installWatchmedo()
}

func installWatchmedo() {
	message := installer.InstallWatchmedo()
	fmt.Println(message)
}

var protobufCmd = &cobra.Command{
	Use:   "protobuf",
	Short: "fer install protobuf",
	Long:  "This subcommand to install protobuf",
	Run:   installWatchmedoCmd,
}

func installProtobufCmd(cmd *cobra.Command, args []string) {
	installProtobuf()
}

func installProtobuf() {
	message := installer.ProtobufInstaller()
	fmt.Println(message)
}
