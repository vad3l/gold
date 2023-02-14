package main

import (
	"image/color"
	"fmt"

	."GUI/Scene"
	."GUI/Widgets"
	."GUI/Utilities"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MenuScene struct {
	biglabel	*Label
}

func NewMenuScene () *MenuScene {

	biglabel := NewLabel("To Show example push Arrow key \n\n					 				A	Z	E",Point{ 0,0 },)
	biglabel.SetFunction(
		func(g *SceneManager) {
			fmt.Println("label")
		})

	biglabel.SetColor(color.RGBA{0xff,0xff,0xff,0xff})
	biglabel.SetFont("../../data/font/TTF/dogicapixel.ttf")
	biglabel.SetFontSize(50)

	
	return &MenuScene{
		&biglabel,
	}
}

func (m *MenuScene) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0xb9, 0xf3, 0xff, 0xff })
	DrawDebug(screen)

	m.biglabel.Draw(screen)
}

func (m *MenuScene) Update(g *SceneManager) error {
	m.biglabel.Input(g)
    if ebiten.IsKeyPressed(ebiten.Key1) { 
		g.Current_scene = NewExampleOne()
    } else if ebiten.IsKeyPressed(ebiten.Key2) {
        g.Current_scene = NewExampleTwo()
    } else if ebiten.IsKeyPressed(ebiten.Key3) {
        
    }

	return nil
}

func (m *MenuScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	return 1280, 720
}
