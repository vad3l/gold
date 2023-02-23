package GUI

import (
	"fmt"
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Quit_game = errors.New("regular termination")
)


type Scene interface {
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)	
	Update(g *SceneManager) error
}

func DrawTPS(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.ActualTPS()))
}

func DrawFPS(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.2f", ebiten.ActualTPS()),0,15)
}

func DrawDebug(screen *ebiten.Image) {
	DrawTPS(screen)
	DrawFPS(screen)
}