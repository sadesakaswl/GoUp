package main

import (
	"fmt"
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
	if err = exec.Command(filepath.Join(home, ".goup", "go", "bin", "go"), "get", "-u", "golang.org/dl/"+version).Run(); err != nil {
		return err
	}
	fmt.Printf("Saved as %s\n", version)
	if download, _ := getCmd.PersistentFlags().GetBool("download"); download {
		fmt.Println("Downloading and Extracting...")
		return exec.Command(filepath.Join(home, "go", "bin", version), "download").Run()
	}
	return nil
}
