package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func getLatestVersion() (string, error) {
	resp, err := http.Get("https://golang.org/dl/?mode=json")

	if err != nil {
		return "", err
	}
	jsonByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	JSON := new([]interface{})
	err = json.Unmarshal(jsonByte, JSON)
	if err != nil {
		return "", err
	}
	return (*JSON)[0].(map[string]interface{})["version"].(string), nil
}
func versionCheck(text string) (string, error) {
	text = strings.ToLower(text)
	if text == "latest" {
		version, err := getLatestVersion()
		if err != nil {
			return "", err
		}
		return version, nil
	}
	if !strings.Contains(text, "go") {
		return strings.Join([]string{"go", text}, ""), nil
	}
	return text, nil
}
func getInstalledVersion() (string, error) {
	cmd := exec.Command("go", "version")
	versionArray, err := cmd.Output()
	if err != nil {
		return "", err
	}
	version := bytes.Fields(versionArray)[2]
	return string(version), nil
}
func check() (int, error) {
	latest, err := getLatestVersion()
	if err != nil {
		return 0, err
	}
	installed, err := getInstalledVersion()
	if err != nil {
		return 0, err
	}
	if latest == installed {
		fmt.Println("Go is already up to date")
		return 0, nil
	} else if strings.Compare(latest, installed) == 1 {
		fmt.Printf("Go is not up to date - %s\n", installed)
		fmt.Printf("Latest version - %s\n", latest)
		fmt.Printf("For upgrade: %s upgrade\n", os.Args[0])
		return 1, nil
	}
	fmt.Printf("Beta version found - %s\n", installed)
	fmt.Printf("To switch latest channel: %s update\n", os.Args[0])
	fmt.Printf("To get other beta version: %s install [VERSION]\n", os.Args[0])
	return -1, nil
}
