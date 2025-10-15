package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	gfx "github.com/vad3l/gold/library/graphics"
)

type Widget interface {
	Draw(screen *ebiten.Image)
	Input()
}

func hover(pCursor, size, position gfx.Point, img *ebiten.Image, scale float64) bool {
	if (pCursor.X <= (size.X+position.X) && pCursor.X >= position.X) && (pCursor.Y <= (size.Y+position.Y) && pCursor.Y >= position.Y) {
		c := img.At(int(pCursor.X/scale)-int(position.X/scale), int(pCursor.Y/scale)-int(position.Y/scale)).(color.RGBA)
		return c.A > 0
	}
	return false
}
