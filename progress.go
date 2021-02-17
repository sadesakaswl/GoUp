package main

import (
	"fmt"
	"strings"
	"time"
)

const progressText = "%s... %s"

//Progress Gives info about process to user
func Progress(status <-chan string, reloadDuration time.Duration) {
	ticker := time.NewTicker(reloadDuration)
	operation := <-status
	keys := []string{"|", "/", "-", `\`}
	for {
		for i := 0; i < len(keys); i++ {
			select {
			case <-ticker.C:
				fmt.Printf(progressText+"\r", operation, keys[i])
			case operation = <-status:
				if operation == "Exit" {
					fmt.Printf(strings.Repeat(" ", len(fmt.Sprintf(progressText, operation, keys[i]))*2) + "\r")
					ticker.Stop()
					return
				}
				if operation == "Pause" {
					operation = <-status
				}
				fmt.Printf(strings.Repeat(" ", len(fmt.Sprintf(progressText, operation, keys[i]))*2) + "\r")
			}
		}
	}
}
func initProgress(duration time.Duration) (statusChan chan string, reloadDuration time.Duration) {
	statusChan = make(chan string)
	return statusChan, duration
}
func closeProgress(statusChan chan<- string) {
	statusChan <- "Exit"
}
func pauseProgress(statusChan chan<- string) {
	statusChan <- "Pause"
}
