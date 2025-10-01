package main

import (
	"fmt"
	"image/color"
	"os"

	. "github.com/vad3l/gold/library/graphics"
	. "github.com/vad3l/gold/library/graphics/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type ExampleOne struct {
	playButton     *Button
	settingsButton *Button
	quitButton     *Button
}

func NewExampleOne() *ExampleOne {

	buttonWidth := 300.0
	buttonHeight := 100.0
	width, _ := ebiten.WindowSize()
	widthWindow := float64(width)/2 - buttonWidth/2

	playButton := NewButton(Point{buttonWidth, buttonHeight}, Point{widthWindow, 100}, "Play")
	settingsButton := NewButton(Point{buttonWidth, buttonHeight}, Point{widthWindow, 300}, "Settings")

	quitButton := NewButton(Point{buttonWidth, buttonHeight}, Point{widthWindow, 500}, "Quit")

	return &ExampleOne{
		&playButton,
		&settingsButton,
		&quitButton,
	}
}

func (m *ExampleOne) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})
	DrawDebug(screen)

	m.playButton.Draw(screen)
	m.settingsButton.Draw(screen)
	m.quitButton.Draw(screen)
}

func (m *ExampleOne) Update(g *SceneManager) error {
	m.playButton.Input()
	m.settingsButton.Input()
	m.quitButton.Input()

	if m.playButton.Execute {
		fmt.Println("play")
		m.playButton.Execute = false
	}

	if m.settingsButton.Execute {
		g.Current_scene = NewSettingsScene()
		m.settingsButton.Execute = false
	}

	if m.quitButton.Execute {
		os.Exit(0)
		m.quitButton.Execute = false
	}

	chooseScene(g)
	return nil
}

func (m *ExampleOne) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}
