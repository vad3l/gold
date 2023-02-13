package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720) 
	ebiten.SetMaxTPS(200)
	g := &SceneManager{
		current_scene: NewMenuScene(),
	}

	if err := ebiten.RunGame(g); err != nil {
		if err != quit_game {
			panic(err)
		}
	}
}

