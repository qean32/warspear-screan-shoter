package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kbinani/screenshot"
)

func getUserChoice() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Действие число от 1 до 3:")
		fmt.Println("  1 - Запуск скрипта")
		fmt.Println("  2 - Перезаход в атр.")
		fmt.Println("  3 - Калибровка (разработка)")
		fmt.Println("  4 - Тест отдельного скрина (разработка)")
		fmt.Println("  5 - Скрины (разработка)")
		fmt.Println("  6 - Выход")
		fmt.Println("  6 - Переключить монитор")
		fmt.Print("Ваш выбор: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: Введите корректное число.")
			continue
		}

		if num < 1 || num > 7 {
			fmt.Println("Ошибка: Введите допустимое число")
			continue
		}

		return num
	}
}

func takeScreenshot(save bool) {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		log.Println("Не найдено активных дисплеев")
		return
	}

	for i := 0; i < n; i++ {
		if i == 1 {
			continue
		}

		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Printf("Ошибка захвата экрана %d: %v", i, err)
			continue
		}

		if save {
			saveScrean(img, i, i)
		} else {
			analyze(img)
		}
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

func saveScrean(img *image.RGBA, index int, i int) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("screenshots-warspear-shoter/display_%d_%s_%d.png", i, timestamp, index)

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Ошибка создания файла: %v", err)
		return
	}
	png.Encode(file, img)
	file.Close()

	fmt.Printf("Сохранено: %s\n", filename)
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	if m > 0 {
		return fmt.Sprintf("%dм %02dс", m, s)
	}
	return fmt.Sprintf("%dс", s)
}

func calibrationTest(img image.Image) bool {
	r8, g8, b8 := pixelAt(img, coordinare{605, 13})

	return colorMatch(r8, g8, b8, 255, 255, 255, 10)
}
