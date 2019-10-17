package deploy

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var devList = []string{"dev-a", "dev-b", "dev-c", "dev-d", "dev-e"}
var beta = "beta"
var preview = "preview"
var stable = "stable"

const (
	deployshell  = "deploy.sh"
	deployscript = `#!/bin/bash -e

readonly BETA=0
readonly STABLE=1
readonly PREVIEW=2
readonly DEV=3
readonly V4=4
readonly MDEV=5
readonly DEVA=6
readonly DEVB=7
readonly SRE=8

generate_version() {
  format="v$(date +%Y%m%d.%s)"
  
  case "$1" in
    0)
      version="beta-$format"
    ;;
    1)
      version="stable-$format"
    ;;
    2)
      version="preview-$format"
    ;;
    3)
      version="$2-$format"
    ;;
    4)
      version="v4-$format"
    ;;
    5)
      version="$2-$format"
    ;;
    6)
      version="deva-$format"
    ;;
    7)
      version="devb-$format"
    ;;
    8)
      version="sre-$format"
    ;;
  esac
  
  echo "Version $version generated."
}

create_tag() {
  read tag_description
  echo Tag description: $tag_description

  git tag -a $version -m "$tag_description"
}

push_tag() {
  echo "Pushing version $version to repository..."

  git push origin $version

  echo "Tag $version succesfully pushed to repository."
}

release_version() {
  echo "Releasing version $version..."

  generate_version $1
  create_tag
  push_tag

  echo "Version release completed."
}

release_mdev() {
  echo "Releasing version mdev ... "
  
  generate_version $1 $2
  create_tag
  push_tag
}

release_dev() {
  echo "Releasing version dev ... "
  
  generate_version $1 $2
  create_tag
  push_tag
}

while getopts "r:s:t:" option; do
  case "${option}" in
    r)
      release_type=${OPTARG}

      case "$release_type" in
        beta)
          release_version $BETA
        ;;
        stable)
          release_version $STABLE
        ;;
        preview)
          release_version $PREVIEW
        ;;
        v4)
          release_version $V4
        ;;
        dev-a)
          release_version $DEVA
        ;;
        dev-b)
          release_version $DEVB
        ;;
        sre)
          release_version $SRE
        ;;
      esac
    ;;
    s)
      release_mdev $MDEV $OPTARG
    ;;
    t)
      release_dev $DEV $OPTARG
    ;;
  esac
done
`
)

type (
	//Deploy define deploy
	Deploy interface {
		Run(target string)
	}
	deploy struct{}
)

//NewDeploy ...
func NewDeploy() Deploy {
	return &deploy{}
}

//CreateScript to create deploy script
func CreateScript() {
	script := []byte(deployscript)
	err := ioutil.WriteFile(deployshell, script, 0644)
	if err != nil {
		_ = os.Remove(deployshell)
		log.Fatal("fail when create deploy script")
	}
}

//
func (b *deploy) Run(target string) {
	CreateScript()
	target, flag, err := checkTarget(target)
	if err != nil {
		_ = os.Remove(deployshell)
		log.Fatal(err)
	}
	cmd := exec.Command("bash", deployshell, flag, target)
	fmt.Println("Deploying to", target)
	cmd.Stdin = os.Stdin
	fmt.Println("Please input tag description:")
	err = runCommand(cmd)
	if err != nil {
		_ = os.Remove(deployshell)
		log.Fatal("fail when run deploy script, error :", err)
	}
}

func checkTarget(target string) (targetResult, flag string, err error) {
	if strings.Contains(target, "dev") {
		flag = "-t"
		for _, value := range devList {
			if target == value {
				targetResult = value
				return
			}
			if target == beta {
				targetResult = target
				return
			}
		}
	}
	if target == preview || target == stable || target == beta {
		targetResult = target
		flag = "-r"
		return
	}
	err = errors.New(`unknown target.
available target : beta,dev-a,dev-b,dev-c,dev-d,dev-e,preview,stable`)
	return
}

func runCommand(cmd *exec.Cmd) error {
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	fmt.Printf("%s\n", outb.String())
	if err != nil {
		return err
	}
	return nil
}
