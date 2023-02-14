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

type ExampleTwo struct {
	playButton	*Button
	settingsButton	*Button
	quitButton	*Button
}

func NewExampleTwo () *ExampleTwo {

	buttonWidth := 300.0
	buttonHeight := 100.0
	width, _ := ebiten.WindowSize()
	widthWindow := float64(width) /2 - buttonWidth /2 
	 
	playButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,100},"Play",
		func(g *SceneManager) {
			fmt.Println("Play")
		})
	playButton.SetRadius(150)
	playButton.SetColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	playButton.SetColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	playButton.SetColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	playButton.SetColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	playButton.SetFont("../../data/font/TTF/dogicapixel.ttf")
	playButton.SetFontSize(35)

	settingsButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,300},"Settings",
		func(g *SceneManager) {
			fmt.Println("Settings")
		})
	settingsButton.SetRadius(150)
	settingsButton.SetColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	settingsButton.SetColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	settingsButton.SetColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	settingsButton.SetColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	settingsButton.SetFont("../../data/font/TTF/dogicapixel.ttf")
	settingsButton.SetFontSize(35)


	quitButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,500},"Quit",
		func(g *SceneManager) {
			os.Exit(0)
		})
	quitButton.SetRadius(150)
	quitButton.SetColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	quitButton.SetColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	quitButton.SetColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	quitButton.SetColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	quitButton.SetFont("../../data/font/TTF/dogicapixel.ttf")
	quitButton.SetFontSize(35)



	return &ExampleTwo{
		&playButton,
		&settingsButton,
		&quitButton,
	}
}

func (m *ExampleTwo) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0xb9, 0xf3, 0xff, 0xff })
	DrawDebug(screen)
	
	m.playButton.Draw(screen)
	m.settingsButton.Draw(screen)
	m.quitButton.Draw(screen)
}

func (m *ExampleTwo) Update(g *SceneManager) error {
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

func (m *ExampleTwo) Layout (outsideWidth, outsideHeight int) (int, int) {	
	return 1280, 720
}
