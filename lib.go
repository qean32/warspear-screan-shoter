package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func script() {
	fmt.Println("скрипт запушен. перенесите фокус на игру")
	run(func() { takeScreenshot(false) }, 1)
}

func calibration() {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		log.Println("Не найдено активных дисплеев")
		return
	}

	bounds := screenshot.GetDisplayBounds(monitor)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Printf("Ошибка захвата экрана %d: %v", 1, err)
		return
	}

	res := calibrationTest(img)

	if res {
		fmt.Println("соответствует ", monitor, " монитор")
	} else {
		fmt.Println("не соответствует ", monitor, " монитор")
	}
}

func reEntry() {
	fmt.Println("в разработке!")
}

func screenshots() {
	if err := os.MkdirAll("screenshots-warspear-shoter", 0755); err != nil {
		log.Fatal(err)
	}

	run(func() { takeScreenshot(true) }, 2)
}

func testFrame() {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		log.Println("Не найдено активных дисплеев")
		return
	}

	bounds := screenshot.GetDisplayBounds(monitor)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Printf("Ошибка захвата экрана %d: %v", monitor, err)
		return
	}

	analyze(img)
}

func run(fn func(), interval int) {
	const (
		logInterval   = 20 * time.Second
		totalDuration = 118 * time.Second
	)
	fnInterval := time.Duration(interval) * time.Second

	start := time.Now()

	ticker := time.NewTicker(fnInterval)
	defer ticker.Stop()

	logTicker := time.NewTicker(logInterval)
	defer logTicker.Stop()

	stopTimer := time.NewTimer(totalDuration)
	defer stopTimer.Stop()

	fn()
	for {
		select {
		case <-ticker.C:
			fn()
		case <-logTicker.C:
			remaining := totalDuration - time.Since(start)
			if remaining < 0 {
				remaining = 0
			}
			fmt.Printf("До остановки скрипта осталось: %s\n", formatDuration(remaining))
		case <-stopTimer.C:
			fmt.Println("Стоп!")
			return
		}
	}
}

func toggle() {
	if monitor == 1 {
		monitor = 0
	} else {
		monitor = 1
	}
}

func screen() {
	bounds := screenshot.GetDisplayBounds(monitor)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Printf("Ошибка захвата экрана %d: %v", monitor, err)
		return
	}

	saveScrean(img, monitor, monitor)
}
