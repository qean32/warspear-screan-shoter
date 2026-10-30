package main

import (
	"log"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	takeScreenshot()
	for range ticker.C {
		takeScreenshot()
	}
}

func takeScreenshot() {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		log.Println("Не найдено активных дисплеев")
		return
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}

		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Printf("Ошибка захвата экрана %d: %v", i, err)
			continue
		}
		analyzeAndAct(img)
	}
}
