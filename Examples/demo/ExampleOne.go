package main

import (
	"image/color"
	"fmt"
	"os"

	."GUI/Scene"
	."GUI/Widgets"
	."GUI/Utilities"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ExampleOne struct {
	playButton	*Button
	settingsButton	*Button
	quitButton	*Button
}

func NewExampleOne () *ExampleOne {

	buttonWidth := 300.0
	buttonHeight := 100.0
	width, _ := ebiten.WindowSize()
	widthWindow := float64(width) /2 - buttonWidth /2 
	 
	playButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,100},"Play",
		func(g *SceneManager) {
			fmt.Println("Play")
		})

	settingsButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,300},"Settings",
		func(g *SceneManager) {
			fmt.Println("Settings")
		})


	quitButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,500},"Quit",
		func(g *SceneManager) {
			os.Exit(0)
		})



	return &ExampleOne{
		&playButton,
		&settingsButton,
		&quitButton,
	}
}

func (m *ExampleOne) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0xb9, 0xf3, 0xff, 0xff })
	DrawDebug(screen)
	
	m.playButton.Draw(screen)
	m.settingsButton.Draw(screen)
	m.quitButton.Draw(screen)
}

func (m *ExampleOne) Update(g *SceneManager) error {
	m.playButton.Input(g)
	m.settingsButton.Input(g)
	m.quitButton.Input(g)

	if ebiten.IsKeyPressed(ebiten.Key1) { 
		g.Current_scene = NewExampleOne()
    } else if ebiten.IsKeyPressed(ebiten.Key2) {
        g.Current_scene = NewExampleTwo()
    } else if ebiten.IsKeyPressed(ebiten.Key3) {
        
    }
	return nil
}

func (m *ExampleOne) Layout (outsideWidth, outsideHeight int) (int, int) {	
	return 1280, 720
}
