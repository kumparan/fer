package cmd

import (
	"fmt"

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
	message := installer.InstallGoUtils("protoc-gen-go", config.ProtobufInstallerURL)
	message = installer.InstallGoUtils("mockgen", config.MockgenInstallerURL)
	message = installer.InstallGoUtils("richgo", config.RichgoInstallerURL)
	message = installer.InstallGoUtils("golint", config.GolintInstallerURL)
	message = installer.InstallGoUtils("git-chglog", config.ChangeLogInstallerURL)
	message = installer.CheckedInstallerPath("make")
	message = installer.InstallWatchmedo()
	fmt.Printf("%+v", message)

}

func init() {
	installCmd.AddCommand(goUtilsCmd)
	installCmd.AddCommand(watchmedoCmd)
	RootCmd.AddCommand(installCmd)
}

var goUtilsCmd = &cobra.Command{
	Use:   "goutils",
	Short: "fer install goutils",
	Long:  "This subcommand to install git go utils like ",
	Run:   installGoUtilsCmd,
}

func installGoUtilsCmd(_ *cobra.Command, _ []string) {
	installGoUtils()
}

func installGoUtils() {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	message := installer.InstallGoUtils("protoc-gen-go", config.ProtobufInstallerURL)
	message = installer.InstallGoUtils("mockgen", config.MockgenInstallerURL)
	message = installer.InstallGoUtils("richgo", config.RichgoInstallerURL)
	message = installer.InstallGoUtils("golint", config.GolintInstallerURL)
	message = installer.InstallGoUtils("git-chglog", config.ChangeLogInstallerURL)
	message = installer.CheckedInstallerPath("make")
	fmt.Printf("%+v", message)
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
