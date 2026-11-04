package main

import "os"

var monitor = 1

func main() {
	for {
		choice := getUserChoice()

		switch choice {
		case 1:
			script()
		case 2:
			reEntry()
		case 3:
			calibration()
		case 4:
			testFrame()
		case 5:
			screenshots()
		case 6:
			os.Exit(1)
		case 7:
			toggle()
		case 8:
			screen()
		}
	}
}

// спасибо адмиралу Вовке за помощь в тестировании!
