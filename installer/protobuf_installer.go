package installer

import (
	"fmt"
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

func protobufDownloadInstaller() string {
	downloadURL := config.ProtobufLinuxInstallerURL
	if runtime.GOOS == "darwin" {
		downloadURL = config.ProtobufOSXInstallerURL
	}
	cmd := exec.Command("curl", "-OL", downloadURL)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return errorDownload
	}
	return successDownload
}

// ProtobufInstaller :nodoc:
func ProtobufInstaller() string {
	message := protobufDownloadInstaller()
	protocZip := config.ProtocZipLinux
	if runtime.GOOS == "darwin" {
		protocZip = config.ProtocZipOSX
	}
	if message == successDownload {
		cmdUnzipToBinProtocPath := exec.Command("sudo", "unzip", "-o", protocZip, "-d", "/usr/local", "bin/protoc")
		err := cmdUnzipToBinProtocPath.Run()
		if err != nil {
			fmt.Println(err)
			return errorUnzip
		}
		cmdProtocPath := exec.Command("sudo", "unzip", "-o", protocZip, "-d", "/usr/local", "'include/*'")
		err = cmdProtocPath.Run()
		if err != nil {
			fmt.Println(err)
			return errorUnzip
		}
		cmdRemoveProtocZip := exec.Command("rm", "-f", protocZip)
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
