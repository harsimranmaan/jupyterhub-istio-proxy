package main

import "fmt"

var (
	version   = "dev"
	commit    = "HEAD"
	goVersion = "Unknown"
)

func getVersionInfo() string {
	return fmt.Sprintf("Version: %s Commit: %s GoVersion: %s", version, commit, goVersion)
}