package main

import (
	"fmt"
	"io/ioutil"
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
	if version == "all" {
		directories, err := ioutil.ReadDir(filepath.Join(home, "sdk"))
		if err != nil {
			return err
		}
		os.RemoveAll(filepath.Join(home, "sdk"))
		for _, v := range directories {
			if err = os.Remove(filepath.Join(home, "go", "bin", v.Name()+extension)); err != nil {
				return err
			}
			fmt.Printf("Deleted %s\n", v.Name())
		}
		return nil
	}
	os.RemoveAll(filepath.Join(home, "sdk", version))
	if err = os.Remove(filepath.Join(home, "go", "bin", version+extension)); err != nil {
		return err
	}
	fmt.Printf("Deleted %s\n", version)
	return nil
}
