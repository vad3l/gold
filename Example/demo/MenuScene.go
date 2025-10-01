package main

import (
	"fmt"
	"image/color"

	. "github.com/vad3l/gold/library/graphics"
	. "github.com/vad3l/gold/library/graphics/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuScene struct {
	biglabel *Label
}

func NewMenuScene() *MenuScene {

	biglabel := NewLabel("To Show example push Arrow key \n\n					 				A	Z	E", Point{0, 0})
	biglabel.SetFunction(
		func(g *SceneManager) {
			fmt.Println("label")
		})

	biglabel.SetColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	biglabel.SetFont("data/dogicapixel.ttf")
	biglabel.SetFontSize(50)

	return &MenuScene{
		&biglabel,
	}
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})
	DrawDebug(screen)

	m.biglabel.Draw(screen)
}

func (m *MenuScene) Update(g *SceneManager) error {
	m.biglabel.Input(g)
	chooseScene(g)

	return nil
}

func (m *MenuScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}
