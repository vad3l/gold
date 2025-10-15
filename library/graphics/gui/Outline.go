package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	gfx "github.com/vad3l/gold/library/graphics"
	"golang.org/x/image/font/basicfont"
)

type Outline struct {
	Position gfx.Point
	Size     gfx.Point
	Text     string
}

func (c *Outline) Draw(screen *ebiten.Image) {
	// modern panel: shadow, fill, border, centered title
	shadowOffset := 6.0
	// shadow
	ebitenutil.DrawRect(screen, c.Position.X+shadowOffset, c.Position.Y+shadowOffset, c.Size.X, c.Size.Y, color.RGBA{0, 0, 0, 40})
	// panel background
	panelColor := color.RGBA{240, 248, 255, 255} // aliceblue-ish
	ebitenutil.DrawRect(screen, c.Position.X, c.Position.Y, c.Size.X, c.Size.Y, panelColor)
	// border
	borderColor := color.RGBA{200, 100, 120, 180}
	ebitenutil.DrawRect(screen, c.Position.X, c.Position.Y, c.Size.X, 2, borderColor)
	ebitenutil.DrawRect(screen, c.Position.X, c.Position.Y+c.Size.Y-2, c.Size.X, 2, borderColor)
	ebitenutil.DrawRect(screen, c.Position.X, c.Position.Y, 2, c.Size.Y, borderColor)
	ebitenutil.DrawRect(screen, c.Position.X+c.Size.X-2, c.Position.Y, 2, c.Size.Y, borderColor)

	// title centered at top
	if c.Text != "" {
		bounds := text.BoundString(basicfont.Face7x13, c.Text)
		textW := float64(bounds.Max.X - bounds.Min.X)
		titleX := c.Position.X + (c.Size.X / 2) - (textW / 2)
		titleY := c.Position.Y + 14
		text.Draw(screen, c.Text, basicfont.Face7x13, int(titleX), int(titleY), color.RGBA{255, 255, 255, 255})
	}
}

func (c *Outline) Input() {
	return
}
