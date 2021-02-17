package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func createGoUpDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(home, ".goup"), 0755)
}
func writeEnvFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(home, ".goup", "env"), []byte(unixPathFile), 0644)
}
func removeGoUpDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(home, ".goup"))
}
