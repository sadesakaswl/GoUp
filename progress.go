package main

import (
	"fmt"
	"time"
)

//Progress Gives info about process to user
func Progress(status <-chan string, reloadDuration time.Duration) {
	ticker := time.NewTicker(reloadDuration)
	operation := <-status
	keys := []string{"|", "/", "-", `\`}
	for {
		for i := 0; i < len(keys); i++ {
			select {
			case <-ticker.C:
				fmt.Printf("%s... %s\r", operation, keys[i])
			case operation = <-status:
				if operation == "Exit" {
					fmt.Println()
					ticker.Stop()
					return
				}
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
