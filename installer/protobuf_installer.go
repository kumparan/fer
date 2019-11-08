package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/kumparan/fer/config"
)

var (
	errorDownload   = "fail downloaded protobuf"
	errorUnzip      = "fail unzip protobuf"
	errorInstall    = "fail installeded protobuf"
	successDownload = "success downloaded protobuf"
)

func protobufDownloadInstaller() (filePath, message string) {
	fmt.Println("Download protobuf")
	downloadURL := config.ProtobufLinuxInstallerURL
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "darwin" {
		downloadURL = config.ProtobufOSXInstallerURL
	}
	err = DownloadFile(filepath.Join(tmp, config.ProtocZipFileName), downloadURL)
	if err != nil {
		fmt.Println(err)
		return "", errorDownload
	}
	ProgressBar(100)
	fmt.Println(successDownload)
	return tmp, successDownload
}

// ProtobufInstaller :nodoc:
func ProtobufInstaller() {
	fmt.Println("Installing protobuf")
	tmp, message := protobufDownloadInstaller()
	protobufZipFile := config.ProtocZipFileName
	filePath := filepath.Join(tmp, protobufZipFile)
	if runtime.GOOS == "darwin" {
		protobufZipFile = config.ProtocZipFileName
	}
	if message == successDownload {
		fmt.Println("Configure protobuf")
		_, err := Unzip(filePath, "/usr/local/bin/")
		if err != nil {
			ProgressBar(1)
			log.Fatal(err)
			fmt.Println(errorUnzip)
			os.Exit(1)
		}

		_, err = Unzip(filePath, "include/*")
		if err != nil {
			ProgressBar(1)
			log.Fatal(err)
			fmt.Println(errorUnzip)
			os.Exit(1)
		}

		cmdRemoveProtocZip := exec.Command("rm", "-f", filePath)
		err = cmdRemoveProtocZip.Run()
		if err != nil {
			ProgressBar(1)
			log.Fatal(err)
			fmt.Println(errorUnzip)
			os.Exit(1)
		}
		checkedMessage := CheckedInstallerPath("protoc")
		ProgressBar(100)
		fmt.Println(checkedMessage)
		os.Exit(0)
	}
	ProgressBar(1)
	fmt.Println(message)
	os.Exit(1)
}
