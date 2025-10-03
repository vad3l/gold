package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	. "github.com/vad3l/gold/library/graphics"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetTPS(200)
	g := &SceneManager{}
	g.AddScene(NewMenuScene())
	g.AddScene(NewExampleOne())
	g.AddScene(NewExampleTwo())
	g.AddScene(NewExampleThree())
	g.AddScene(NewSettingsScene())

	if err := ebiten.RunGame(g); err != nil {
		if err != Quit_game {
			panic(err)
		}
	}
}

func chooseScene(g *SceneManager) {
	if inpututil.IsKeyJustReleased(ebiten.Key1) {
		g.ChangeScene("ExampleOne")
	} else if inpututil.IsKeyJustReleased(ebiten.Key2) {
		g.ChangeScene("ExampleTwo")
	} else if inpututil.IsKeyJustReleased(ebiten.Key3) {
		g.ChangeScene("ExampleThree")
	}
}
