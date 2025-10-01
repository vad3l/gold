package main

import (
	"fmt"
	"image/color"

	. "github.com/vad3l/gui"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SettingsScene struct {
	bigoutline *Outline
	biglabel   *Label
	//bigcheckbox	*Checkbox
	backButton  *Button
	basicButton *Button
}

func NewSettingsScene() *SettingsScene {

	/*bigcheckbox := Checkbox{
		0,
		[]string{ "niggure", "nighga" },
		[]Point{ Point{200, 200}, Point{300, 200} },
		Point{50, 50},
		func(g *SceneManager, i int) {
			fmt.Println("bigCheckBox")
		},
	}*/

	bigoutline := Outline{
		Point{500.0, 500.0},
		Point{100, 200},
		"bigoutline",
	}

	biglabel := NewLabel("To Show example push Arrow key \n\n					 				<--     -->", Point{0, 0})
	biglabel.SetFunction(
		func(g *SceneManager) {
			fmt.Println("label")
		})

	biglabel.SetColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	biglabel.SetFont("data/dogicapixel.ttf")
	biglabel.SetFontSize(50)

	backButton := NewButton(Point{150, 50}, Point{50, 50}, "BACK")
	backButton.SetRadius(150)
	backButton.SetColor(color.RGBA{0xb5, 0xf1, 0xcc, 0xff})
	backButton.SetColorHover(color.RGBA{0xe5, 0xfd, 0xd1, 0xff})
	backButton.SetColorText(color.RGBA{0xFF, 0xAA, 0xCF, 0xff})
	backButton.SetColorTextHover(color.RGBA{0xFF, 0xAA, 0xCF, 0xff})
	backButton.SetFont("data/dogicapixel.ttf")
	backButton.SetFontSize(20)

	basicButton := NewButton(Point{100, 200}, Point{800, 500}, "Basic")
	return &SettingsScene{
		&bigoutline,
		&biglabel,
		//&bigcheckbox,
		&backButton,
		&basicButton,
	}
}

func (m *SettingsScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})
	m.backButton.Draw(screen)
	//m.bigcheckbox.Draw(screen)
	m.biglabel.Draw(screen)
	m.bigoutline.Draw(screen)
	m.basicButton.Draw(screen)
}

func (m *SettingsScene) Update(g *SceneManager) error {
	m.backButton.Input()
	m.biglabel.Input(g)
	//m.bigcheckbox.Input(g)
	m.basicButton.Input()

	if m.backButton.Execute {
		g.Current_scene = NewMenuScene()
		m.backButton.Execute = false
	}
	return nil
}

func (m *SettingsScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}
