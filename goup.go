package main

import (
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
func removeGoUpDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(home, ".goup"))
}
