package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const unixPathFile = `#!/bin/sh
# goup 
case ":${PATH}:" in
    *:"$HOME/.goup/go/bin":*)
        ;;
    *)
        export PATH="$HOME/.goup/go/bin:$PATH"
        ;;
esac
case ":${PATH}:" in
    *:"$HOME/go/bin":*)
        ;;
    *)
        export PATH="$HOME/go/bin:$PATH"
        ;;
esac
export GOROOT="$HOME/.goup/go"
export GOPATH="$HOME/go"
`
const profileLine = `source "$HOME/.goup/env"`
const windowsPath = `%USERPROFILE%\.goup\go\bin;%USERPROFILE%\go\bin`
const windowsGoPath = `%USERPROFILE%\go`
const windowsGoRootPath = `%USERPROFILE%\.goup\go`

func checkPathUnix() (bool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}
	profile, err := ioutil.ReadFile(filepath.Join(home, ".profile"))
	if err != nil {
		return false, err
	}
	return bytes.Contains(profile, []byte(profileLine)), nil
}
func addPathUnix() error {
	contains, err := checkPathUnix()
	if err != nil {
		return err
	}
	if contains {
		return nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	profileFile, err := os.OpenFile(filepath.Join(home, ".profile"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer profileFile.Close()
	if err != nil {
		return err
	}
	_, err = profileFile.WriteString(profileLine)
	if err != nil {
		return err
	}
	return nil
}
func removePathUnix() error {
	contains, err := checkPathUnix()
	if err != nil {
		return err
	}
	if !contains {
		return nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	profile, err := ioutil.ReadFile(filepath.Join(home, ".profile"))
	if err != nil {
		return err
	}
	profileArray := bytes.Split(profile, []byte(profileLine))
	profile = bytes.Join(profileArray, nil)
	err = ioutil.WriteFile(filepath.Join(home, ".profile"), profile, 0644)
	if err != nil {
		return err
	}
	return nil
}
func checkPathWindows() (bool, error) {
	path, err := exec.Command("REG", "QUERY", `HKCU\Environment`, "/v", "Path").Output()
	if err != nil {
		return false, err
	}
	return bytes.Contains(path, []byte(windowsPath)), nil
}
func checkGoPathWindows() bool {
	path, err := exec.Command("REG", "QUERY", `HKCU\Environment`, "/v", "GOPATH").Output()
	if err != nil {
		return false
	}
	return bytes.Contains(path, []byte(windowsGoPath))
}
func addPathWindows() error {
	contains, err := checkPathWindows()
	if err != nil {
		return err
	}
	if contains {
		if err := removeWindowsPath(); err != nil {
			return err
		}
	}
	path, err := getWindowsPath()
	if err != nil {
		return err
	}
	path = bytes.Join([][]byte{path, []byte(windowsPath)}, []byte(";"))
	exec.Command("REG", "ADD", `HKCU\Environment`, "/v", "Path", "/t", "REG_EXPAND_SZ", "/d", string(path), "/f").Output()
	return nil
}
func addGoPathWindows() error {
	contains := checkGoPathWindows()
	if contains {
		if err := removeWindowsGoPath(); err != nil {
			return nil
		}
	}
	exec.Command("REG", "ADD", `HKCU\Environment`, "/v", "GOPATH", "/t", "REG_EXPAND_SZ", "/d", windowsGoPath, "/f").Output()
	exec.Command("REG", "ADD", `HKCU\Environment`, "/v", "GOROOT", "/t", "REG_EXPAND_SZ", "/d", windowsGoRootPath, "/f").Output()
	return nil
}
func getWindowsPath() ([]byte, error) {
	output, err := exec.Command("REG", "QUERY", `HKCU\Environment`, "/v", "Path").Output()
	if err != nil {
		return nil, err
	}
	outputArray := bytes.Fields(output)[3:]
	output = bytes.Join(outputArray, []byte(" "))
	return output, nil
}
func removeWindowsPath() error {
	path, err := getWindowsPath()
	if err != nil {
		return err
	}
	if !bytes.Contains(path, []byte(windowsPath)) {
		return nil
	}
	pathArray := bytes.Split(path, []byte(windowsPath))
	found := len(pathArray) - 1
	path = bytes.Join(pathArray, nil)
	path = path[:len(path)-found]
	_, err = exec.Command("REG", "ADD", `HKCU\Environment`, "/v", "Path", "/t", "REG_EXPAND_SZ", "/d", string(path), "/f").Output()
	if err != nil {
		return err
	}
	return nil
}
func removeWindowsGoPath() error {
	exec.Command("REG", "DELETE", "HKCU\\Environment", "/v", "GOPATH", "/f").Output()
	exec.Command("REG", "DELETE", "HKCU\\Environment", "/v", "GOROOT", "/f").Output()
	return nil
}
func chmodUnix() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filepath.Walk(filepath.Join(home, ".goup"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(path, 0777)
	})
	filepath.Walk(filepath.Join(home, "go"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(path, 0777)
	})
	return nil
}
