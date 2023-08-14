package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	Current_scene	Scene
}

func (g *SceneManager) Draw(screen *ebiten.Image) {
	g.Current_scene.Draw(screen)
}

func (g *SceneManager) Update() error {
	return g.Current_scene.Update(g)	
}

func (g *SceneManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Current_scene.Layout(1280, 720)
}
