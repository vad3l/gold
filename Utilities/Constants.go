package Utilities

import (
	"errors"
	"image/color"
)

var (
	quit_game = errors.New("regular termination")
	foreground = color.RGBA{0xff, 0x00, 0x00, 0xff}
	background = color.RGBA{0x00, 0xff, 0x00, 0xff}
)

type Point struct {
	X float64
	Y float64
}

