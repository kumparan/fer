package deployment

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	//AvailableTargets define to info the command
	AvailableTargets = "available targets: dev,dev-a,dev-b,dev-c,dev-d,dev-e,dev-f,staging,prod"
)

var targetMap = map[string]string{
	"dev":     "beta",
	"dev-a":   "dev-a",
	"dev-b":   "dev-b",
	"dev-c":   "dev-c",
	"dev-d":   "dev-d",
	"dev-e":   "dev-e",
	"dev-f":   "dev-f",
	"staging": "preview",
	"prod":    "stable",
}

type (
	//Deployment define deployment
	Deployment interface {
		Run(target string)
	}
	deployment struct{}
)

//NewDeployment ...
func NewDeployment() Deployment {
	return &deployment{}
}

//
func (d *deployment) Run(target string) {
	reader := bufio.NewReader(os.Stdin)
	target, err := checkTarget(target)
	if err != nil {
		log.Fatal("unknown target.\n", AvailableTargets)
	}
	tagTime := CreateTagTime()
	tag := target + "-" + tagTime
	fmt.Println("Releasing verion to |", target, "|")
	fmt.Println("Version " + tag)
	fmt.Println("Please Input Tag Description: ")
	desc, _ := reader.ReadString('\n')
	fmt.Print("Tag description: " + desc)
	fmt.Println("Pushing version", tag, "to repository ...")
	createTagCmd := exec.Command("git", "tag", "-a", tag, "-m", desc)
	err = runCommand(createTagCmd)
	if err != nil {
		log.Fatal(err)
	}
	pushTagCmd := exec.Command("git", "push", "origin", tag)
	err = runCommand(pushTagCmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done push", tag)

}

//CreateTagTime create unique tag time
func CreateTagTime() string {
	now := time.Now()
	second := strconv.Itoa(int(now.UnixNano()))[:10]
	date := strings.ReplaceAll(now.Format("2006-01-02"), "-", "")
	return fmt.Sprintf("v%s.%s", date, second)
}

func checkTarget(target string) (targetResult string, err error) {
	for key, value := range targetMap {
		if key == target {
			targetResult = value
			return
		}
	}
	err = errors.New("unknown target")
	return
}

func runCommand(cmd *exec.Cmd) error {
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println(outb.String(), errb.String())
	return nil
}
