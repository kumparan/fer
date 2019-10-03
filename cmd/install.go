package cmd

import (
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
	installer.InstallProtobuf()
	installer.InstallMockgen()
	installer.InstallRichgo()
	installer.InstallGoLint()
	installer.InstallChangelog()
	installer.CheckExistenceOfChangelog()
	installer.CheckChangelogVersion()
	installer.CheckExistenceOfMake()
	installer.InstallWatchmedo()
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
	installer.InstallProtobuf()
	installer.InstallMockgen()
	installer.InstallRichgo()
	installer.InstallGoLint()
	installer.InstallChangelog()
	installer.CheckExistenceOfChangelog()
	installer.CheckChangelogVersion()
	installer.CheckExistenceOfMake()
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
	installer.InstallWatchmedo()
}
