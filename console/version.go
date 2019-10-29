package console

import (
	"encoding/json"
	"fmt"
	"github.com/blang/semver"
	"github.com/kumparan/fer/cache"
	"github.com/kumparan/fer/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const (
	lastVerCacheKey = "last_version_cache_key"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print fer version",
	Long:  `print version of fer`,
	Run:   printVersion,
}

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Println(config.Version)
}

type githubRelease struct {
	TagName string `json:"tag_name"`
}

func getFerLatestVersion() (ver string, err error) {
	resp, err := http.Get(config.ReleaseURL)
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var jsonData []githubRelease
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		return
	}
	ver = jsonData[0].TagName
	return
}

//getFerLatestVersionCached
func getFerLatestVersionCached() (ver string, err error) {
	pathDir := filepath.Join(GetConfigDir(), "cache")
	c, err := cache.New(pathDir)
	if err != nil {
		return
	}

	v, err := c.GetOrSetFunc(lastVerCacheKey, func() (interface{}, error) {
		latestVer, err2 := getFerLatestVersion()
		if err2 != nil {
			return nil, err2
		}
		latestVerB := []byte(latestVer)
		return latestVerB, nil
	})

	if err != nil {
		return
	}

	if ver, ok := v.([]byte); ok {
		return string(ver), nil
	}

	return "", fmt.Errorf("unable to parse version from cache: %v", v)
}

func checkVersion() {
	latestVer, err := getFerLatestVersionCached()
	if err != nil {
		PrintWarn("Error getting latest version: %s\n", err.Error())
	}
	currentVersion, err := semver.Make(config.Version[1:])
	if err != nil {
		PrintWarn("Error parsing current version: %s\n", err.Error())
		return
	}
	latestVersion, err := semver.Make(latestVer[1:])
	if err != nil {
		PrintWarn("Error parsing latest version: %s\n", err.Error())
	}

	if currentVersion.LT(latestVersion) {
		PrintWarn("Your installed fer is out of date (%s), please update to the latest version: %s!\n", config.Version, latestVer)
	}
}
