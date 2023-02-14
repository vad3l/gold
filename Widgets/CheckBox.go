package Widgets

import (
	"image/color"
	."GUI/Utilities"
	."GUI/Scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type Checkbox struct {
	Size	Point

	position	Point
	colorCheckBox	color.RGBA
	colorCheckBoxChecked	color.RGBA
	borderSize	float64
	checked	bool
}

func NewCheckBox (size, position Point) Checkbox {

	return Checkbox{
		size,
		position,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		10.0,
		false,
	}
}

func (c *Checkbox) Draw (screen *ebiten.Image) {

}

func (c *Checkbox) Input (g *SceneManager) {

}

func (c *Checkbox) hover (pSize ,pRect, pCursor Point) bool  {
	return ((pCursor.X <= (pSize.X + pRect.X) && pCursor.X >= pRect.X) && (pCursor.Y <= (pSize.Y + pRect.Y) && pCursor.Y >= pRect.Y))
}