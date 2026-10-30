package main

import (
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

// checkPoint описывает одну проверяемую точку:
// где смотреть, какой цвет ожидается и какую клавишу нажать,
// если цвет НЕ совпал.
type checkPoint struct {
	coordinare coordinare
	line       string
	expected   rgb
	keyCode    uint16 // виртуальный код клавиши Windows (VK_...)
	tolerance  uint8
}

// Заранее подготовленный массив из 5 точек.
// У каждой точки свой ожидаемый цвет и своя клавиша.
var checkPoints = [5]checkPoint{
	{line: "ПЕРВАЯ", coordinare: coordinare{800, 300}, expected: rgb{112, 63, 13}, keyCode: 0x31, tolerance: 10},   // 1
	{line: "ВТОРАЯ", coordinare: coordinare{860, 300}, expected: rgb{97, 57, 18}, keyCode: 0x32, tolerance: 10},    // 2
	{line: "ТРЕТЬЯ", coordinare: coordinare{925, 300}, expected: rgb{92, 48, 13}, keyCode: 0x33, tolerance: 10},    // 3
	{line: "ЧЕТВЕРТАЯ", coordinare: coordinare{995, 300}, expected: rgb{97, 47, 11}, keyCode: 0x34, tolerance: 10}, // 4
	{line: "ПЯТАЯ", coordinare: coordinare{1065, 300}, expected: rgb{107, 61, 9}, keyCode: 0x35, tolerance: 10},    // 5
}

// analyzeAndAct проверяет цвет в каждой точке. Если цвет НЕ совпадает
// с ожидаемым для этой точки — нажимает соответствующую клавишу.
func analyzeAndAct(img image.Image) {
	for _, cp := range checkPoints {
		pixelColor := img.At(cp.coordinare.x, cp.coordinare.y)

		r, g, b, _ := pixelColor.RGBA()
		r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

		if !colorMatch(r8, g8, b8, cp.expected.r, cp.expected.g, cp.expected.b, cp.tolerance) {
			log.Printf("Линия: %v, точка (%d,%d) (rgb(%d, %d, %d) - rgb(%d, %d, %d))",
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

// Вспомогательная функция для сравнения с допуском
func colorMatch(r1, g1, b1, r2, g2, b2, tol uint8) bool {
	diff := func(a, b uint8) int {
		if a > b {
			return int(a - b)
		}
		return int(b - a)
	}
	return diff(r1, r2) <= int(tol) && diff(g1, g2) <= int(tol) && diff(b1, b2) <= int(tol)
}
