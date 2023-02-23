package GUI

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Widget interface {
	draw(g *SceneManager)
	input(g *SceneManager)

}


func hover (pCursor, size, position Point, img	*ebiten.Image, scale float64) bool  {
	if ((pCursor.X <= (size.X + position.X) && pCursor.X >= position.X) && (pCursor.Y <= (size.Y + position.Y) && pCursor.Y >= position.Y)){
		c := img.At(int(pCursor.X/scale) - int(position.X /scale), int(pCursor.Y/scale)- int(position.Y /scale) ).(color.RGBA)	
		if c.A > 0 {
			return true
		}
		return false
	}
	return false
}
