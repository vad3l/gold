package graphics

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	scenes        []Scene
	current_scene Scene
}

func (g *SceneManager) AddScene(scene Scene) {
	g.scenes = append(g.scenes, scene)
	if g.current_scene == nil {
		g.current_scene = scene
	}
}

func (g *SceneManager) Draw(screen *ebiten.Image) {
	if g.current_scene != nil {
		g.current_scene.Draw(screen)
	}
}

func (g *SceneManager) Update() error {
	if g.current_scene != nil {
		return g.current_scene.Update(g)
	}
	return nil
}

func (g *SceneManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	if g.current_scene != nil {
		return g.current_scene.Layout(outsideWidth, outsideHeight)
	}
	return outsideWidth, outsideHeight
}

func (g *SceneManager) ChangeScene(name string) error {
	for _, scene := range g.scenes {
		if scene.Name() == name {
			g.current_scene = scene
			return nil
		}
	}
	return errors.New("scene not found: " + name)
}
