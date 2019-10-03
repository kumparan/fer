package installer

import (
	"fmt"
	"os"
	"os/exec"
)

// InstallChangelog :nodoc:
func InstallChangelog() {
	cmd := exec.Command("go", "get", "-u", "github.com/git-chglog/git-chglog/cmd/git-chglog")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// InstallProtobuf :nodoc:
func InstallProtobuf() {
	cmd := exec.Command("go", "get", "-u", "github.com/golang/protobuf/protoc-gen-go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// InstallMockgen :nodoc:
func InstallMockgen() {
	cmd := exec.Command("go", "get", "-u", "github.com/golang/mock/mockgen")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// InstallRichgo :nodoc:
func InstallRichgo() {
	cmd := exec.Command("go", "get", "-u", "github.com/kyoh86/richgo")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// InstallGoLint :nodoc:
func InstallGoLint() {
	cmd := exec.Command("go", "get", "-u", "github.com/golangci/golangci-lint/cmd/golangci-lint")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
