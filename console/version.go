package console

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"gopkg.in/djherbis/fscache.v0"
	"io"
	"log"
	"time"

	"github.com/kumparan/fer/config"
	"github.com/spf13/cobra"
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
	TagName            string  `json:"tag_name"`
}

//getGithubReleaseLatest
func getGithubReleaseLatest() (string) {
	c, err := fscache.New(".cache-version", 0755, 2*time.Hour)
	if err != nil{
		log.Fatal(err.Error())
	}

	r, w, err := c.Get("last_version")
	if err != nil {
		log.Fatal(err)
	}

	if w == nil {
		buf := new(bytes.Buffer)
		io.Copy(buf, r)
		r.Close()
		ver := buf.String()
		return ver
	}

	resp, err := http.Get(config.APIFerGithubReleaseURL)
	if err != nil{
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return ""
	}

	var jsonData []githubRelease
	err = json.Unmarshal([]byte(body), &jsonData) // here!
	if err != nil {
		fmt.Println("error unmarshal")
		return ""
	}
	var version = jsonData[0].TagName
	w.Write([]byte(version))
	w.Close()
	return version
}

func checkVersion() {
	if getGithubReleaseLatest() != config.Version{
		fmt.Println("your fer is out of date, please updated your fer")
	}
}