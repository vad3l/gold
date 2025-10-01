package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	. "github.com/vad3l/gui"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetMaxTPS(200)
	g := &SceneManager{
		Current_scene: NewMenuScene(),
	}

	if err := ebiten.RunGame(g); err != nil {
		if err != Quit_game {
			panic(err)
		}
	}
}

func chooseScene(g *SceneManager) {
	if inpututil.IsKeyJustReleased(ebiten.Key1) {
		g.Current_scene = NewExampleOne()
	} else if inpututil.IsKeyJustReleased(ebiten.Key2) {
		g.Current_scene = NewExampleTwo()
	} else if inpututil.IsKeyJustReleased(ebiten.Key3) {
		g.Current_scene = NewExampleThree()
	}
}
