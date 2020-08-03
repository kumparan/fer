package console

import (
	"os"

	"github.com/kumparan/fer/config"
	"github.com/kumparan/fer/installer"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install dependencies for your project",
	Long:  "install what do you need for backend contributed",
}

func init() {
	installCmd.AddCommand(installAllCmd)
	installCmd.AddCommand(goUtilsCmd)
	installCmd.AddCommand(protocGenCmd)
	installCmd.AddCommand(mockgenCmd)
	installCmd.AddCommand(richgoCmd)
	installCmd.AddCommand(golintCmd)
	installCmd.AddCommand(chglogCmd)
	installCmd.AddCommand(protobufCmd)
	installCmd.AddCommand(moddCmd)
}

var installAllCmd = &cobra.Command{
	Use:   "all",
	Short: "This subcommand to install all dependencies",
	Long:  "install all dependencies for contributing to backend projects",
	Run:   installAll,
}

func installAll(_ *cobra.Command, _ []string) {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	installer.InstallGoUtils("protoc-gen-go", config.ProtobufInstallerURL)
	installer.InstallGoUtils("mockgen", config.MockgenInstallerURL)
	installer.InstallGoUtils("richgo", config.RichgoInstallerURL)
	installer.InstallGoUtils("golint", config.GolintInstallerURL)
	installer.InstallGoUtils("git-chglog", config.ChangeLogInstallerURL)
	installer.ProtobufInstaller()
	installer.InstallModd()
	os.Exit(0)
}

var goUtilsCmd = &cobra.Command{
	Use:   "goutils",
	Short: "This subcommand to install all go utils",
	Long:  "To install all goutils, your golang version must %s or latest",
	Run:   installGoUtilsCmd,
}

func installGoUtilsCmd(_ *cobra.Command, _ []string) {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	installer.InstallGoUtils("protoc-gen-go", config.ProtobufInstallerURL)
	installer.InstallGoUtils("mockgen", config.MockgenInstallerURL)
	installer.InstallGoUtils("richgo", config.RichgoInstallerURL)
	installer.InstallGoUtils("golint", config.GolintInstallerURL)
	installer.InstallGoUtils("git-chglog", config.ChangeLogInstallerURL)
	os.Exit(0)
}

var protocGenCmd = &cobra.Command{
	Use:   "protoc-gen",
	Short: "This subcommand to install protoc generator",
	Long:  "Go version must be " + config.GoVersion + " or latest",
	Run:   installProtocGenCmd,
}

func installProtocGenCmd(_ *cobra.Command, _ []string) {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	installer.InstallGoUtils("protoc-gen-go", config.ProtobufInstallerURL)
	os.Exit(0)
}

var mockgenCmd = &cobra.Command{
	Use:   "mockgen",
	Short: "This subcommand to install mock generator",
	Long:  "GoMock is a mocking framework for the Go programming language. It integrates well with Go's built-in testing package, but can be used in other contexts too.",
	Run:   installMockgenCmd,
}

func installMockgenCmd(_ *cobra.Command, _ []string) {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	installer.InstallGoUtils("mockgen", config.MockgenInstallerURL)
	os.Exit(0)
}

var richgoCmd = &cobra.Command{
	Use:   "richgo",
	Short: "This subcommand to install richgo",
	Long:  "Rich-Go will enrich go test outputs with text decorations.",
	Run:   installRichgoCmd,
}

func installRichgoCmd(_ *cobra.Command, _ []string) {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	installer.InstallGoUtils("richgo", config.MockgenInstallerURL)
	os.Exit(0)
}

var golintCmd = &cobra.Command{
	Use:   "golint",
	Short: "This subcommand to install GolangCI-Lint",
	Long:  "GolangCI-Lint is a linters aggregator. It's easy to integrate and use, has nice output and has a minimum number of false positives. It supports go modules.",
	Run:   installGolintCmd,
}

func installGolintCmd(_ *cobra.Command, _ []string) {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	installer.InstallGoUtils("golint", config.MockgenInstallerURL)
	os.Exit(0)
}

var chglogCmd = &cobra.Command{
	Use:   "git-chglog",
	Short: "This subcommand to install git-changelog",
	Long:  "Changelog generator implemented in Go (Golang)",
	Run:   installChglogCmd,
}

func installChglogCmd(_ *cobra.Command, _ []string) {
	installer.CheckExistenceOfGolang()
	installer.CheckGolangVersion()
	installer.InstallGoUtils("git-chglog", config.MockgenInstallerURL)
	os.Exit(0)
}

var protobufCmd = &cobra.Command{
	Use:   "protobuf",
	Short: "This subcommand to install protobuf",
	Long:  "Protocol Buffers (a.k.a., protobuf) are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data.",
	Run:   installProtobufCmd,
}

func installProtobufCmd(cmd *cobra.Command, args []string) {
	installer.ProtobufInstaller()
	os.Exit(0)
}

var moddCmd = &cobra.Command{
	Use:   "modd",
	Short: "This subcommand to install modd",
	Long:  "Modd is a developer tool that triggers commands and manages daemons in response to filesystem changes. https://github.com/cortesi/modd",
	Run:   installModdCmd,
}

func installModdCmd(cmd *cobra.Command, args []string) {
	installer.InstallModd()
	os.Exit(0)
}
