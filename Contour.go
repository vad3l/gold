package main

import (
	"golang.org/x/image/font/basicfont"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type Contour struct {
	position Point
	size Point
	text string
}

func (c *Contour) Draw (screen *ebiten.Image) {
	/*
		0,0 -> 0,y
		0,y -> x,y
		x,0 -> x,y
		0,t -> x,0
	*/
	
	text.Draw(screen, c.text, basicfont.Face7x13, int(c.position.x + 5), int(c.position.y), foreground)
	ebitenutil.DrawLine(screen, c.position.x, c.position.y - 13, c.position.x, c.position.y + c.size.y, foreground)
	ebitenutil.DrawLine(screen, c.position.x, c.position.y + c.size.y, c.position.x + c.size.x, c.position.y + c.size.y, foreground)
	ebitenutil.DrawLine(screen, c.position.x + c.size.x, c.position.y, c.position.x + c.size.x, c.position.y + c.size.y, foreground)
	ebitenutil.DrawLine(screen, c.position.x + float64((len(c.text) * 7) + 10), c.position.y, c.position.x + c.size.x, c.position.y, foreground)
}

func (c *Contour) Input (g *SceneManager) {
	return
}
