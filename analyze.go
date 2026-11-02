package main

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/dacapoday/sendinput"
)

var checkPoints = [5]checkPoint{
	{
		line: "ПЕРВАЯ", keyCode: 0x31, tolerance: 70,
		variants: []variant{
			{coordinare{791, 420}, rgb{18, 166, 113}},
			{coordinare{791, 490}, rgb{18, 166, 113}},
		},
	},
	{
		line: "ВТОРАЯ", keyCode: 0x32, tolerance: 70,
		variants: []variant{
			{coordinare{856, 420}, rgb{18, 166, 113}},
			{coordinare{856, 490}, rgb{18, 166, 113}},
		},
	},
	{
		line: "ТРЕТЬЯ", keyCode: 0x33, tolerance: 70,
		variants: []variant{
			{coordinare{927, 420}, rgb{18, 166, 113}},
			{coordinare{927, 490}, rgb{18, 166, 113}},
		},
	},
	{
		line: "ЧЕТВЕРТАЯ", keyCode: 0x34, tolerance: 70,
		variants: []variant{
			{coordinare{992, 420}, rgb{18, 166, 113}},
			{coordinare{992, 490}, rgb{18, 166, 113}},
		},
	},
	{
		line: "ПЯТАЯ", keyCode: 0x35, tolerance: 70,
		variants: []variant{
			{coordinare{1080, 420}, rgb{18, 166, 113}},
			{coordinare{1080, 490}, rgb{18, 166, 113}},
		},
	},
}

const observer = true

func analyze(img image.Image) {
	for _, cp := range checkPoints {
		for i, v := range cp.variants {
			r8, g8, b8 := pixelAt(img, v.coordinare)
			matched := colorMatch(r8, g8, b8, v.expected.r, v.expected.g, v.expected.b, cp.tolerance)

			if !matched && !observer {
				fmt.Printf("[%s] вариант #%d точка (%d,%d): rgb(%d,%d,%d) ожидается rgb(%d,%d,%d) tol=%d\n",
					cp.line, i+1,
					v.coordinare.x, v.coordinare.y,
					r8, g8, b8,
					v.expected.r, v.expected.g, v.expected.b,
					cp.tolerance,
				)
			}

			if matched && observer {
				log.Printf("Линия: %v, точка (%d,%d) — цвет совпал",
					cp.line, v.coordinare.x, v.coordinare.y)

				if false {
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
	}
}
