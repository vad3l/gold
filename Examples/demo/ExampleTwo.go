package main

import (
	"image/color"
	."Framework"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
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
	 
	playButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,100},"Play")
	playButton.SetRadius(150)
	playButton.SetColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	playButton.SetColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	playButton.SetColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	playButton.SetColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	playButton.SetFont("./data/dogicapixel.ttf")
	playButton.SetFontSize(35)

	settingsButton := playButton
	settingsButton.SetText("Settings")
	settingsButton.SetPosition(Point{widthWindow,300})

	quitButton := playButton
	quitButton.SetText("Quit")
	quitButton.SetPosition(Point{widthWindow,500})

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
	m.playButton.Input()
	m.settingsButton.Input()
	m.quitButton.Input()

	if (m.playButton.Execute) {
		fmt.Println("play")
		m.playButton.Execute = false
	}

	if (m.settingsButton.Execute) {
		g.Current_scene = NewSettingsScene()
		m.settingsButton.Execute = false
	}

	if (m.quitButton.Execute) {
		os.Exit(0)
		m.quitButton.Execute = false
	}

	chooseScene(g)
	return nil
}

func (m *ExampleTwo) Layout (outsideWidth, outsideHeight int) (int, int) {	
	return 1280, 720
}
