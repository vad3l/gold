package main

import (
	."GUI/Scene"

	"github.com/hajimehoshi/ebiten/v2"
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

