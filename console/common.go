package console

import (
	"github.com/gookit/color"
	"github.com/kumparan/fer/config"
	"log"
	"os/user"
	"path/filepath"
)

// PrintError :nodoc:
func PrintError(format string, args ...interface{}) {
	color.Error.Printf(format, args...)
}

// PrintInfo :nodoc:
func PrintInfo(format string, args ...interface{}) {
	color.Info.Printf(format, args...)
}

// PrintWarn :nodoc:
func PrintWarn(format string, args ...interface{}) {
	color.Warn.Printf(format, args...)
}

// GetConfigDir :nodoc:
func GetConfigDir() string {
	return filepath.Join(GetHomeDir(), config.ConfigDir)
}

// GetHomeDir :nodoc:
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
