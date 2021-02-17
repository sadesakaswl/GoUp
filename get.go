package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func get(version string) error {
	version, err := versionCheck(version)
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return exec.Command(filepath.Join(home, ".goup", "go", "bin", "go"), "get", "-u", "golang.org/dl/"+version).Run()
}
