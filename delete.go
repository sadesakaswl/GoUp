package main

import (
	"os"
	"path/filepath"
	"runtime"
)

func delete(version string) error {
	version, err := versionCheck(version)
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	var extension string
	if runtime.GOOS == "windows" {
		extension = ".exe"
	} else {
		extension = ""
	}
	return os.Remove(filepath.Join(home, "go", "bin", version+extension))
}
