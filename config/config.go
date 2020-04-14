package config

import (
	"os"
	"path/filepath"
)

// FerConfigPath :nodoc:
func FerConfigPath() string {
	return filepath.Join(os.Getenv("HOME"), ".ferconfig.json")
}
