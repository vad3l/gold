package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	current_scene	Scene
}

func (g *SceneManager) Draw(screen *ebiten.Image) {
	g.current_scene.Draw(screen)
}

func (g *SceneManager) Update() error {
	return g.current_scene.Update(g)	
}

func (g *SceneManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.current_scene.Layout(1280, 720)
}
