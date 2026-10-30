package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	// Создаём папку для скриншотов
	if err := os.MkdirAll("screenshots", 0755); err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	counter := 1

	// Делаем первый скриншот сразу
	takeScreenshot(counter)
	counter++

	for range ticker.C {
		takeScreenshot(counter)
		counter++
	}
}

func takeScreenshot(index int) {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		log.Println("Не найдено активных дисплеев")
		return
	}

	for i := 0; i < n; i++ {
		// Пропускаем основной монитор (индекс 0)
		if i == 0 {
			continue
		}

		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Printf("Ошибка захвата экрана %d: %v", i, err)
			continue
		}

		timestamp := time.Now().Format("2006-01-02_15-04-05")
		filename := fmt.Sprintf("screenshots/display_%d_%s_%d.png", i, timestamp, index)

		file, err := os.Create(filename)
		if err != nil {
			log.Printf("Ошибка создания файла: %v", err)
			continue
		}
		png.Encode(file, img)
		file.Close()

		fmt.Printf("Сохранено: %s\n", filename)
	}
}
