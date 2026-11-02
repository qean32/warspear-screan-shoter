package main

import (
	"image"
	"log"

	"github.com/kbinani/screenshot"
)

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
		analyze(img)
	}
}

func colorMatch(r1, g1, b1, r2, g2, b2, tol uint8) bool {
	diff := func(a, b uint8) int {
		if a > b {
			return int(a - b)
		}
		return int(b - a)
	}
	return diff(r1, r2) <= int(tol) && diff(g1, g2) <= int(tol) && diff(b1, b2) <= int(tol)
}

func pixelAt(img image.Image, c coordinare) (uint8, uint8, uint8) {
	r, g, b, _ := img.At(c.x, c.y).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}
