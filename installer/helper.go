package installer

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	pb "github.com/cheggaaa/pb/v3"

	"github.com/kumparan/fer/config"

	version "github.com/hashicorp/go-version"
)

// CheckExistenceOfGolang :nodoc:
func CheckExistenceOfGolang() {
	cmdGetGolangLocation := exec.Command("which", "go")
	err := cmdGetGolangLocation.Run()
	if err != nil {
		fmt.Println("You should install golang first")
		os.Exit(1)
	}
}

// CheckGolangVersion :nodoc:
func CheckGolangVersion() {
	cmdGetGolangVersion, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	var goLocalversion = string(cmdGetGolangVersion)
	var regexVersion, _ = regexp.Compile(`(\d+\.\d+\.\d+)`)
	v1, err := version.NewVersion(config.GoVersion)
	if err != nil {
		log.Fatal(err)
	}
	v2, err := version.NewVersion(regexVersion.FindString(goLocalversion))
	if err != nil {
		log.Fatal(err)
	}
	if v2.LessThan(v1) {
		fmt.Printf("Go version must be %s or latest\n", config.GoVersion)
		os.Exit(1)
	}
}

// CheckedInstallerPath :nodoc:
func CheckedInstallerPath(installer string) string {
	cmdGetInstallerPath := exec.Command("which", installer)
	err := cmdGetInstallerPath.Run()
	if err != nil {
		return "fail installed " + installer
	}
	return "Success installed " + installer
}

// DownloadFile :nodoc:
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Unzip :nodoc:
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip.
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

// ProgressBar :nodoc:
func ProgressBar(stopProgress int) {
	count := stopProgress

	// start bar from 'full' template
	bar := pb.Full.Start(100)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond * 100)
	}
	bar.Finish()
}

// CreateChangelogYAML :nodoc:
func CreateChangelogYAML(style, repositoryURL string) string {
	var configurationText = `style: ` + style + `
	template: CHANGELOG.tpl.md
	info:
	  title: CHANGELOG
	  repository_url:` + repositoryURL + `
	options:
	  commits:
		 filters:
		   Type:
			 - feature
			 - bugfix
			 - hotfix
			 - refactor
			 - test
			 - misc
	  commit_groups:
		 title_maps:
		   feature: New Features
		   bugfix: Fixes
		   hotfix: Fixes
		   refactor: Code Improvements
		   test: Test Improvements
		   misc: Other Improvements
	  header:
		pattern: "^(\\w*)\\:\\s(.*)$"
		pattern_maps:
		  - Type
		  - Subject
	  notes:
		keywords:
		  - BREAKING CHANGE
	`

	return configurationText
}

// CreateChangelogMD :nodoc:
func CreateChangelogMD() string {
	var changelogmd = `{{ if .Versions -}}

{{ if .Unreleased.CommitGroups -}}
{{ range .Unreleased.CommitGroups -}}
### {{ .Title }}
{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ range .Versions }}
<a name="{{ .Tag.Name }}"></a>
## {{ if .Tag.Previous }}[{{ .Tag.Name }}]{{ else }}{{ .Tag.Name }}{{ end }} - {{ datetime "2006-01-02" .Tag.Date }}
{{ range .CommitGroups -}}
### {{ .Title }}
{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
### {{ .Title }}
{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{- if .Versions }}
[Unreleased]: {{ .Info.RepositoryURL }}/compare/{{ $latest := index .Versions 0 }}{{ $latest.Tag.Name }}...HEAD
{{ range .Versions -}}
{{ if .Tag.Previous -}}
[{{ .Tag.Name }}]: {{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}
{{ end -}}
{{ end -}}
{{ end -}}
`

	return changelogmd

}
