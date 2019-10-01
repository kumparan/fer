package cmd

import (
	"context"

	"github.com/kumparan/fer/repository"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install all dependencies you are need",
	Long:  "This subcommand start the test",
}

var chgLogCmd = &cobra.Command{
	Use:   "chglog",
	Short: "install all dependencies to generate changelog",
	Long:  "This subcommand to install git change-log",
	RunE:  installchglogCmd,
}

func init() {
	installCmd.AddCommand(chgLogCmd)
	RootCmd.AddCommand(installCmd)
}

func installchglogCmd(cmd *cobra.Command, args []string) error {
	return installChglog(context.TODO())
}

func installChglog(ctx context.Context) error {
	repository.CheckExistenceOfGolang()
	repository.CheckGolangVersion()
	repository.InstalltheChangelog()
	repository.CheckExistenceOfChangelog()
	repository.CheckChangelogVersion()

	return nil
}
