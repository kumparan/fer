package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
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
	downloadURL := config.ProtobufLinuxInstallerURL
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "darwin" {
		downloadURL = config.ProtobufOSXInstallerURL
	}
	err = DownloadFile(tmp+"/"+config.ProtocZipFileName, downloadURL)
	if err != nil {
		fmt.Println(err)
		return "", errorDownload
	}
	return tmp, successDownload
}

// ProtobufInstaller :nodoc:
func ProtobufInstaller() string {
	tmp, message := protobufDownloadInstaller()
	protobufZipFile := config.ProtocZipFileName
	filePath := tmp + "/" + protobufZipFile
	if runtime.GOOS == "darwin" {
		protobufZipFile = config.ProtocZipFileName
	}
	if message == successDownload {
		_, err := Unzip(filePath, "/usr/local/bin/")
		if err != nil {
			log.Fatal(err)
			return errorUnzip
		}

		_, err = Unzip(filePath, "include/*")
		if err != nil {
			log.Fatal(err)
			return errorUnzip
		}

		cmdRemoveProtocZip := exec.Command("rm", "-f", filePath)
		err = cmdRemoveProtocZip.Run()
		if err != nil {
			fmt.Println(err)
			return errorUnzip
		}
		checkedMessage := CheckedInstallerPath("protoc")
		return checkedMessage
	}
	return errorInstall
}
