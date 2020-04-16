package gocek

import "time"

// Module :nodoc:
type Module struct {
	Path    string    `json:"Path"`
	Version string    `json:"Version"`
	Time    time.Time `json:"Time"`
	Update  struct {
		Path    string    `json:"Path"`
		Version string    `json:"Version"`
		Time    time.Time `json:"Time"`
	} `json:"Update"`
	Indirect bool   `json:"Indirect"`
	Dir      string `json:"Dir"`
	GoMod    string `json:"GoMod"`
}

// SimpleModule :nodoc:
type SimpleModule struct {
	Path        string `json:"path"`
	Version     string `json:"version"`
	NextVersion string `json:"next_version"`
}
