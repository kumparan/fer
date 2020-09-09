package gocek

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

const (
	_defaultMaxQueueSize = int(10)
	timeLayout           = "2006-01-02"
)

// ModuleChecker :nodoc:
type ModuleChecker struct {
	RootDir      string
	ServicesDirs []string
}

// CheckCWD call checkCWD() and print into stdout
func (mc *ModuleChecker) CheckCWD() {
	modlist, err := checkCWD()
	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Path", "Version", "Next Version"})

	for _, v := range modlist {
		table.Append([]string{v.Path, v.Version, v.NextVersion})
	}
	table.Render() // Send output
}

// Checks check module projects and save the json file
func (mc *ModuleChecker) Checks() {
	now := time.Now()
	folder := now.Format(timeLayout)
	savePath := path.Join(mc.RootDir, folder)
	err := os.MkdirAll(savePath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range mc.ServicesDirs {
		os.Chdir(dir)
		modules, err := checkCWD()
		if err != nil {
			continue
		}

		sdir := strings.Split(dir, "/")
		err = mc.save(savePath, sdir[len(sdir)-1], modules)
		if err != nil {
			log.Error(err)
			return
		}
		os.Chdir(mc.RootDir)
	}
}

// save modules as json
func (mc *ModuleChecker) save(savePath, modName string, modules []*SimpleModule) error {
	fileName := fmt.Sprintf("%s.json", modName)
	f, err := os.Create(path.Join(savePath, fileName))
	if err != nil {
		log.Error(err)
		return err
	}
	defer f.Close()

	log.Infof("saving %s", fileName)

	j, err := json.MarshalIndent(modules, " ", "	")
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(j))
	return err
}

// checkCWD check current working directory
func checkCWD() (modules []*SimpleModule, err error) {
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

	// group the module per QueueSize,
	// so each group can be conccurently execute
	for _, m := range mods {
		queue[count] = append(queue[count], m)
		if len(queue[count]) == _defaultMaxQueueSize {
			count++
		}
	}

	modsCh := make(chan *Module, len(mods))
	for _, q := range queue {
		wg := sync.WaitGroup{}

		wg.Add(len(q))
		for _, m := range q {
			go func(m string) {
				defer wg.Done()

				mod, err := findModuleUpdate(m)
				if err != nil {
					log.Error("findModuleUpdate: ", err)
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
		log.Error("findDirectModList: ", err)
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
