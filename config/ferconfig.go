package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

var cfg FerConfig
var once sync.Once

// GocekConfig :nodoc:
type GocekConfig struct {
	ProjectDirs   []string `json:"project_directories"`
	SaveOutputDir string   `json:"output_save_folder"`
}

// FerConfig :nodoc:
type FerConfig struct {
	Gocek GocekConfig `json:"gocek"`
}

func init() {
	once.Do(func() {
		loadCfg()
	})
}

func loadCfg() {
	cfgPath := FerConfigPath()
	_, err := os.Stat(cfgPath)
	switch {
	case os.IsNotExist(err):
		f, err := os.Create(cfgPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		cfg.Gocek.ProjectDirs = make([]string, 0)
		bt, _ := json.MarshalIndent(cfg, "", "	")
		_, err = f.Write([]byte(bt))
		if err != nil {
			log.Fatal(err)
		}

	case err == nil:
		b, err := ioutil.ReadFile(FerConfigPath())
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(b, &cfg)
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal(err)
	}
}

// GetFerConfig :nodoc:
func GetFerConfig() FerConfig {
	return cfg
}
