package main

import (
	"fmt"
	"image/color"
	"os"

	. "github.com/vad3l/gold/library/graphics"
	. "github.com/vad3l/gold/library/graphics/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type ExampleThree struct {
	playButton     *SpriteButton
	settingsButton *SpriteButton
	quitButton     *SpriteButton
}

func NewExampleThree() *ExampleThree {

	buttonWidth := 300.0
	width, _ := ebiten.WindowSize()
	widthWindow := float64(width)/2 - buttonWidth/2

	playButton := NewSpriteButton(Point{widthWindow, 100}, "data/play.png", "data/playHover.png", "data/playClicked.png")
	playButton.SetScale(0.2)

	settingsButton := NewSpriteButton(Point{widthWindow, 300}, "data/settings.png", "data/settingsHover.png", "data/settingsClicked.png")
	settingsButton.SetScale(0.2)

	quitButton := NewSpriteButton(Point{widthWindow, 500}, "data/quit.png", "data/quitHover.png", "data/quitClicked.png")
	quitButton.SetScale(0.2)

	return &ExampleThree{
		&playButton,
		&settingsButton,
		&quitButton,
	}
}

func (m *ExampleThree) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})
	DrawDebug(screen)

	m.playButton.Draw(screen)
	m.settingsButton.Draw(screen)
	m.quitButton.Draw(screen)
}

func (m *ExampleThree) Update(g *SceneManager) error {
	m.playButton.Input()
	m.settingsButton.Input()
	m.quitButton.Input()

	if m.playButton.Execute {
		fmt.Println("play")
		m.playButton.Execute = false
	}

	if m.settingsButton.Execute {
		g.ChangeScene("SettingsScene")
		m.settingsButton.Execute = false
	}

	if m.quitButton.Execute {
		os.Exit(0)
		m.quitButton.Execute = false
	}

	chooseScene(g)
	return nil
}

func (m *ExampleThree) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}

func (m *ExampleThree) Name() string {
	return "ExampleThree"
}
