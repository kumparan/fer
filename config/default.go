package config

import (
	"path/filepath"
	"time"
)

const (
	// GoVersion :nodoc:
	GoVersion = "1.12.7"
	// ChangeLogInstallerURL :nodoc:
	ChangeLogInstallerURL = "github.com/git-chglog/git-chglog/cmd/git-chglog"
	// ProtobufInstallerURL :nodoc:
	ProtobufInstallerURL = "github.com/golang/protobuf/protoc-gen-go"
	// MockgenInstallerURL :nodoc:
	MockgenInstallerURL = "github.com/golang/mock/mockgen"
	// RichgoInstallerURL :nodoc:
	RichgoInstallerURL = "github.com/kyoh86/richgo"
	// GolintInstallerURL :nodoc:
	GolintInstallerURL = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	ProtobufVersion    = "3.7.1"
	// ProtobufOSXInstallerURL :nodoc:
	ProtobufOSXInstallerURL = "https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-osx-x86_64.zip"
	// ProtobufLinuxInstallerURL :nodoc:
	ProtobufLinuxInstallerURL = "https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip"
	// ProtocZipFileName :nodoc:
	ProtocZipFileName = "protobuf.zip"
	// ReleaseURL :nodoc:
	ReleaseURL = "https://api.github.com/repos/kumparan/fer/releases"
	// CacheDirPerm :nodoc:
	CacheDirPerm = 0755
	// CacheTTLFerVersion :nodoc:
	CacheTTLFerVersion = 2 * time.Hour

	// TempDirPrefix :nodoc:
	TempDirPrefix = "fer-"
)

var (
	// ConfigDir :nodoc:
	ConfigDir = filepath.Join(".config", "fer")
)
