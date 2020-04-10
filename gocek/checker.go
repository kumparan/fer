package gocek

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/kumparan/go-utils"
	log "github.com/sirupsen/logrus"
)

const (
	_defaultMaxQueueSize = int(10)
)

// ModuleChecker :nodoc:
type ModuleChecker struct {
}

// Checks check module projects and save the json file
func (mc *ModuleChecker) Checks(dirs []string) {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		os.Chdir(dir)
		modules, err := CheckCWD()
		if err != nil {
			continue
		}

		err = mc.save(rootDir, modules)
		if err != nil {
			log.Error(err)
			return
		}
	}
}

// save modules as json
func (mc *ModuleChecker) save(rootDir string, modules []*SimpleModule) error {
	now := time.Now()
	layout := "2006-01-02"

	dst := fmt.Sprintf("%s/bin/%s.mod.json", rootDir, now.Format(layout))
	f, err := os.Create(dst)
	if err != nil {
		log.Error(err)
		return err
	}
	defer f.Close()

	log.Info("saving file...")

	_, err = f.Write([]byte(utils.Dump(modules)))
	return err
}

// CheckCWD check current working directory
func CheckCWD() (modules []*SimpleModule, err error) {
	modList, err := findDirectModList()
	if err != nil {
		log.Error(err)
		return
	}

	return findAllModuleUpdate(modList)
}

func findAllModuleUpdate(mods []string) (modules []*SimpleModule, err error) {
	queue := make(map[int][]string)
	count := 0

	for _, m := range mods {
		queue[count] = append(queue[count], m)
		if len(queue[count]) == _defaultMaxQueueSize {
			count++
		}
	}

	modsCh := make(chan *Module, len(mods))
	for _, mods := range queue {
		wg := sync.WaitGroup{}

		for _, m := range mods {
			wg.Add(1)
			go func(m string) {
				defer wg.Done()

				mod, err := findModuleUpdate(m)
				if err != nil {
					log.Error(err)
					return
				}

				modsCh <- mod
			}(m)
		}

		wg.Wait()
	}

	close(modsCh)
	for module := range modsCh {
		if module.Update.Version == "" {
			continue
		}
		modules = append(modules, &SimpleModule{
			Path:        module.Path,
			Version:     module.Version,
			NextVersion: module.Update.Version,
		})
	}

	return
}

func findModuleUpdate(modName string) (*Module, error) {
	var err error
	out, err := exec.Command("go", "list", "-m", "-u", "-json", modName).Output()
	if err != nil {
		log.WithField("mod", modName).Error(err)
		return nil, err
	}

	m := Module{}
	if err = json.Unmarshal(out, &m); err != nil {
		log.WithField("out", string(out)).Error(err)
		return nil, err
	}

	return &m, nil
}

func findModList() ([]string, error) {
	out, err := exec.Command("go", "list", "-m", "all").Output()
	if err != nil {
		return nil, err
	}
	splitted := strings.Split(string(out), "\n")

	var list []string
	for _, s := range splitted {
		ss := sanitize(s)
		if ss == "" {
			continue
		}
		list = append(list, ss)
	}

	return list, nil
}

func findDirectModList() ([]string, error) {
	out, err := exec.Command("go", "list", "-m", "-f", `{{ .Path }} | {{ .Indirect }}`, "all").Output()
	if err != nil {
		return nil, err
	}
	splitted := strings.Split(string(out), "\n")

	var list []string
	for _, ss := range splitted {
		if strings.Trim(ss, " ") == "" || strings.Contains(ss, "true") {
			continue
		}

		list = append(list, strings.Split(ss, " | ")[0])
	}

	return list, nil
}

func sanitize(raw string) string {
	clean := strings.Trim(raw, " ")
	return strings.Split(clean, " ")[0]
}
