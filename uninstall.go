package main

import (
	"runtime"
	"time"
)

func uninstallWindows() error {
	ch, dur := initProgress(600 * time.Millisecond)
	go Progress(ch, dur)
	ch <- "Removing GoUp to Path environment"
	err := removeWindowsPath()
	if err != nil {
		closeProgress(ch)
		return err
	}
	ch <- "Removing GoUp folder"
	err = removeWindowsGoPath()
	if err != nil {
		return err
	}
	defer closeProgress(ch)
	return removeGoUpDir()
}
func uninstallUnix() error {
	ch, dur := initProgress(600 * time.Millisecond)
	go Progress(ch, dur)
	ch <- "Removing GoUp to Path environment"
	err := removePathUnix()
	if err != nil {
		closeProgress(ch)
		return err
	}
	ch <- "Removing GoUp folder"
	defer closeProgress(ch)
	return removeGoUpDir()
}
func uninstall() error {
	if runtime.GOOS == "windows" {
		err := uninstallWindows()
		if err != nil {
			return err
		}
		return nil
	}
	return uninstallUnix()
}
