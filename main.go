package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	stopTimer := time.NewTimer(2 * time.Minute)
	defer stopTimer.Stop()

	takeScreenshot()
	for {
		select {
		case <-ticker.C:
			takeScreenshot()
		case <-stopTimer.C:
			return
		}
	}
}
