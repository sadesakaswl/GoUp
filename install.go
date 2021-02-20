package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func install2Windows(version string) error {
	ch, dur := initProgress(600 * time.Millisecond)
	go Progress(ch, dur)
	ch <- "Fetching"
	version, err := versionCheck(version)
	path, _ := exec.LookPath("go")
	if path != "" {
		latest, err := check()
		if err != nil {
			return err
		}
		if latestVersion, _ := getLatestVersion(); latestVersion == version && latest == 0 {
			return nil
		}
		if latest == 1 {
			uninstall()
		}
	}
	pauseProgress(ch)
	fmt.Printf("Installing - %s\n", version)
	resp, err := http.Get(fmt.Sprintf("https://dl.google.com/go/%s.windows-%s.zip", version, runtime.GOARCH))
	if err != nil {
		closeProgress(ch)
		return err
	}
	ch <- "Downloading"
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		closeProgress(ch)
		return err
	}
	ch <- "Extracting"
	zipReader, err := zip.NewReader(bytes.NewReader(body), resp.ContentLength)
	if err != nil {
		closeProgress(ch)
	}
	err = createGoUpDir()
	if err != nil {
		closeProgress(ch)
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		closeProgress(ch)
		return err
	}
	for _, zipFile := range zipReader.File {
		path := filepath.Join(home, ".goup", zipFile.Name)
		if zipFile.FileInfo().IsDir() {
			os.MkdirAll(path, zipFile.Mode())
			continue
		}
		fileReader, err := zipFile.Open()
		if err != nil {
			closeProgress(ch)
			return err
		}
		defer fileReader.Close()
		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
		if err != nil {
			closeProgress(ch)
			return err
		}
		if _, err := io.Copy(targetFile, fileReader); err != nil {
			targetFile.Close()
			closeProgress(ch)
			return err
		}
		targetFile.Close()
	}
	ch <- "Adding GoUp to Path Environment"
	addPathWindows()
	addGoPathWindows()
	closeProgress(ch)
	return installGoUp()
}
func install2Unix(version string) error {
	ch, dur := initProgress(600 * time.Millisecond)
	go Progress(ch, dur)
	ch <- "Fetching"
	version, err := versionCheck(version)
	path, _ := exec.LookPath("go")
	if path != "" {
		latest, err := check()
		if err != nil {
			return err
		}
		if latestVersion, _ := getLatestVersion(); latestVersion == version && latest == 0 {
			return nil
		}
		if latest == 1 {
			uninstall()
		}
	}
	pauseProgress(ch)
	fmt.Printf("Installing - %s\n", version)
	resp, err := http.Get(fmt.Sprintf("https://dl.google.com/go/%s.%s-%s.tar.gz", version, runtime.GOOS, runtime.GOARCH))
	if err != nil {
		closeProgress(ch)
		return err
	}
	ch <- "Downloading and Extracting"
	defer resp.Body.Close()
	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		closeProgress(ch)
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		closeProgress(ch)
		return err
	}
	tarReader := tar.NewReader(gzReader)
	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			closeProgress(ch)
			return err
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filepath.Join(home, ".goup", header.Name), 0755); err != nil {
				closeProgress(ch)
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(filepath.Join(home, ".goup", header.Name))
			if err != nil {
				closeProgress(ch)
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				closeProgress(ch)
				outFile.Close()
				return err
			}
			outFile.Close()
		default:
			return fmt.Errorf("Failed on extracting: %v in %s", header.Typeflag, header.Name)
		}
	}
	ch <- "Adding GoUp to Path Environment"
	err = writeEnvFile()
	if err != nil {
		closeProgress(ch)
		return err
	}
	err = addPathUnix()
	if err != nil {
		closeProgress(ch)
		return err
	}
	err = chmodUnix()
	if err != nil {
		closeProgress(ch)
		return err
	}
	closeProgress(ch)
	return installGoUp()
}
func install(version string) error {
	if runtime.GOOS == "windows" {
		err := install2Windows(version)
		return err
	}
	return install2Unix(version)
}
func installGoUp() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	_, err = exec.Command(filepath.Join(home, ".goup", "go", "bin", "go"), "get", "-u", "github.com/sadesakaswl/goup").Output()
	return err
}
