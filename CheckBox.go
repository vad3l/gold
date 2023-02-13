package main

import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Checkbox struct {
	checked int
	label []string
	position []Point
	size Point
	execute func(g *SceneManager, i int)
}

func (c *Checkbox) Draw (screen *ebiten.Image) {
	for i := 0; i < len(c.label); i++ {
		ebitenutil.DrawRect(screen, c.position[i].x, c.position[i].y, c.size.x, c.size.y, foreground)
		if c.checked == i {
			ebitenutil.DrawRect(screen, c.position[i].x + 5, c.position[i].y + 5, c.size.x - 10, c.size.y - 10, background)
		}
		text.Draw(screen, c.label[i], basicfont.Face7x13, int(c.position[i].x + c.size.x + 5), int(c.position[i].y + (c.size.y / 2) + 3), foreground)
	}
}

func (c *Checkbox) Input (g *SceneManager) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pCursor := Point{float64(x),float64(y)}
		for j := 0; j < len(c.label); j++ {
			if in_rect(c.size, c.position[j], pCursor) {
				if c.checked == j {
					c.checked = -1
				} else {
					c.checked = j
				}
				c.execute(g, j)
				break
			}
		}
	}
}

func in_rect (pSize ,pRect, pCursor Point) bool  {
	return ((pCursor.x <= (pSize.x + pRect.x) && pCursor.x >= pRect.x) && (pCursor.y <= (pSize.y + pRect.y) && pCursor.y >= pRect.y))
}