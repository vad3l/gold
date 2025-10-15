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
	slider         *Slider // Ajout du slider
	listView       *ListView
}

func NewExampleOne() *ExampleOne {

	buttonWidth := 300.0
	buttonHeight := 100.0
	width, _ := ebiten.WindowSize()
	widthWindow := float64(width)/2 - buttonWidth/2

	playButton := NewButton(Point{X: buttonWidth, Y: buttonHeight}, Point{X: widthWindow, Y: 100}, "Play")
	settingsButton := NewButton(Point{X: buttonWidth, Y: buttonHeight}, Point{X: widthWindow, Y: 300}, "Settings")
	quitButton := NewButton(Point{X: buttonWidth, Y: buttonHeight}, Point{X: widthWindow, Y: 500}, "Quit")

	// Instanciation du slider
	slider := NewSlider(Point{X: 400, Y: 40}, Point{X: widthWindow, Y: 650}, 0, 100, 50)

	// Create a ListView and populate with buttons
	lv := NewListView(Point{X: 100, Y: 150}, Point{X: 400, Y: 400})
	for i := 0; i < 12; i++ {
		btn := NewButton(Point{X: 360, Y: 40}, Point{X: 0, Y: 0}, fmt.Sprintf("Item %d", i+1))
		// Use a pointer when adding so ListView can modify Execute
		lv.Add(&btn)
	}

	return &ExampleOne{
		&playButton,
		&settingsButton,
		&quitButton,
		&slider, // Ajout du slider
		lv,
	}
}

func (m *ExampleOne) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})
	DrawDebug(screen)

	m.playButton.Draw(screen)
	m.settingsButton.Draw(screen)
	m.quitButton.Draw(screen)
	m.slider.Draw(screen) // Dessine le slider
	if m.listView != nil {
		m.listView.Draw(screen)
	}
}

func (m *ExampleOne) Update(g *SceneManager) error {
	m.playButton.Input()
	m.settingsButton.Input()
	m.quitButton.Input()
	m.slider.Input() // Gère l'input du slider
	if m.listView != nil {
		m.listView.Input()
	}

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

func (m *ExampleOne) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}

func (m *ExampleOne) Name() string {
	return "ExampleOne"
}
