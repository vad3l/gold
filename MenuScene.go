package main

import (
	"image/color"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MenuScene struct {
	playButton	*Button
	settingsButton	*Button
	quitButton	*Button
	spriteButton	*SpriteButton
	pauseButton	*SpriteButton
}

func NewMenuScene () *MenuScene {
	buttonWidth := 300.0
	buttonHeight := 100.0
	width, _ := ebiten.WindowSize()
	widthWindow := float64(width) /2 - buttonWidth /2 
	 
	playButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,100},"Play",
	func(g *SceneManager) {
		fmt.Println("Play")
	})
	playButton.setRadius(150)
	playButton.setColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	playButton.setColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	playButton.setColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	playButton.setColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	playButton.setFont("TTF/dogicapixel.ttf")
	playButton.setFontSize(35)

	settingsButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,300},"Settings",
	func(g *SceneManager) {
		g.current_scene = NewSettingsScene()
	})
	settingsButton.setRadius(150)
	settingsButton.setColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	settingsButton.setColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	settingsButton.setColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	settingsButton.setColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	settingsButton.setFont("TTF/dogicapixel.ttf")
	settingsButton.setFontSize(35)


	quitButton := NewButton(Point{buttonWidth,buttonHeight},Point{widthWindow,500},"Quit",
	func(g *SceneManager) {
		os.Exit(0)
	})
	quitButton.setRadius(150)
	quitButton.setColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	quitButton.setColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	quitButton.setColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	quitButton.setColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	quitButton.setFont("TTF/dogicapixel.ttf")
	quitButton.setFontSize(35)



	spriteButton := NewSpriteButton(Point{530,500},"menuButton.png","menuButtonSelected.png",
	func(g *SceneManager) {
		os.Exit(0)
	})	
	spriteButton.setScale(0.2)
	
	pauseButton := NewSpriteButton(Point{800,300},"pauseUp.png","pausePush.png",
	func(g *SceneManager) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen());
	})	
	pauseButton.setScale(0.2)
	return &MenuScene{
		&playButton,
		&settingsButton,
		&quitButton,
		&spriteButton,
		&pauseButton,
	}
}

func (m *MenuScene) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0xb9, 0xf3, 0xff, 0xff })
	DrawDebug(screen)
	
	//ebitenutil.DrawRect(screen, 100, 100, 200, 150, color.RGBA{0xff, 0x00, 0x00, 0xff})
	//ebitenutil.DrawCircle(screen,200.0,200.0,150.0,color.RGBA{ 0xb9, 0xf3, 0x00, 0xff })
	
	m.playButton.Draw(screen)
	m.settingsButton.Draw(screen)
	//m.quitButton.Draw(screen)
	m.spriteButton.Draw(screen)
	m.pauseButton.Draw(screen)
}

func (m *MenuScene) Update(g *SceneManager) error {
	m.playButton.Input(g)
	m.settingsButton.Input(g)
	//m.quitButton.Input(g)
	m.spriteButton.Input(g)
	m.pauseButton.Input(g)
	return nil
}

func (m *MenuScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	return 1280, 720
}
