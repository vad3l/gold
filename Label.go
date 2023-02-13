package main

import (
        "github.com/hajimehoshi/ebiten/v2"
        "github.com/hajimehoshi/ebiten/v2/inpututil"
        "github.com/hajimehoshi/ebiten/v2/text"
        "golang.org/x/image/font/basicfont"
)

type Label struct {
	text string
	position Point
	execute func(g *SceneManager)
}

func (l *Label) Draw (screen *ebiten.Image) {
	text.Draw(screen, l.text, basicfont.Face7x13, int(l.position.x), int(l.position.y), foreground)
}

func (l *Label) Input (g *SceneManager) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

		x, y := ebiten.CursorPosition()
		size := Point{ float64(7 * len(l.text)), 13 }
		if float64(x) >= l.position.x && float64(x) <= (l.position.x + size.x) && float64(y) >= l.position.y && float64(y) <= (l.position.y + size.y) {
			l.execute(g)
		}
	}
}
