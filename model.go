package main

type coordinare struct {
	x int
	y int
}

type rgb struct {
	r, g, b uint8
}

type variant struct {
	coordinare coordinare
	expected   rgb
}

type checkPoint struct {
	line      string
	keyCode   uint16
	tolerance uint8
	variants  []variant
}
