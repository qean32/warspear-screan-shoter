package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	stopTimer := time.NewTimer(118 * time.Second)
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
