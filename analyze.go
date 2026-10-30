package main

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/dacapoday/sendinput"
)

type coordinare struct {
	x int
	y int
}

type rgb struct {
	r, g, b uint8
}

type checkPoint struct {
	coordinare coordinare
	line       string
	expected   rgb
	keyCode    uint16
	tolerance  uint8
}

var checkPoints = [5]checkPoint{
	{line: "ПЕРВАЯ", coordinare: coordinare{791, 229}, expected: rgb{121, 196, 104}, keyCode: 0x31, tolerance: 70},     // 1
	{line: "ВТОРАЯ", coordinare: coordinare{856, 227}, expected: rgb{120, 190, 95}, keyCode: 0x32, tolerance: 70},     // 2
	{line: "ТРЕТЬЯ", coordinare: coordinare{927, 227}, expected: rgb{72, 159, 96}, keyCode: 0x33, tolerance: 70},     // 3
	{line: "ЧЕТВЕРТАЯ", coordinare: coordinare{992, 229}, expected: rgb{98, 191, 111}, keyCode: 0x34, tolerance: 70}, // 4
	{line: "ПЯТАЯ", coordinare: coordinare{1080, 257}, expected: rgb{112, 186, 90}, keyCode: 0x35, tolerance: 70},     // 5
}

const observer = false

func analyzeAndAct(img image.Image) {
	for _, cp := range checkPoints {
		pixelColor := img.At(cp.coordinare.x, cp.coordinare.y)

		r, g, b, _ := pixelColor.RGBA()
		r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
		if observer {
			fmt.Println(r8, g8, b8, cp.line)
		}

		if colorMatch(r8, g8, b8, cp.expected.r, cp.expected.g, cp.expected.b, cp.tolerance) {
			log.Printf("Линия: %v, точка (%d,%d) — цвет совпал (rgb(%d, %d, %d) ≈ rgb(%d, %d, %d))",
				cp.line,
				cp.coordinare.x, cp.coordinare.y, r8, g8, b8,
				cp.expected.r, cp.expected.g, cp.expected.b)

			if err := sendinput.SendKeyboardInput(sendinput.KeyCode(cp.keyCode), true); err != nil {
				log.Println("Ошибка нажатия:", err)
			}
			time.Sleep(50 * time.Millisecond)
			if err := sendinput.SendKeyboardInput(sendinput.KeyCode(cp.keyCode), false); err != nil {
				log.Println("Ошибка отпускания:", err)
			}
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
