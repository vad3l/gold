package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/vad3l/gold/library/graphics"
	. "github.com/vad3l/gold/library/graphics/gui"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SettingsScene struct {
	bigoutline   *Outline
	biglabel     *Label
	bigcheckbox  *CheckBox
	bigcheckbox2 *CheckBox
	backButton   *Button
	basicButton  *Button
}

func NewSettingsScene() *SettingsScene {

	bigcheckbox := NewCheckBox(Point{30, 30}, Point{100, 400})
	bigcheckbox.SetRadius(8)
	bigcheckbox.SetBorderSize(5)
	bigcheckbox.SetColor(color.RGBA{0, 0, 200, 255})
	bigcheckbox.SetColorChecked(color.RGBA{0, 200, 0, 255})

	bigcheckbox2 := NewCheckBox(Point{30, 30}, Point{150, 450})
	bigcheckbox2.SetRadius(8)
	bigcheckbox2.SetBorderSize(5)
	bigcheckbox2.SetColor(color.RGBA{0, 0, 200, 255})
	bigcheckbox2.SetColorChecked(color.RGBA{0, 200, 0, 255})

	radioGroup := NewRadioGroup()
	bigcheckbox.SetRadioGroup(radioGroup)
	bigcheckbox2.SetRadioGroup(radioGroup)

	radioGroup.Add(&bigcheckbox)
	radioGroup.Add(&bigcheckbox2)

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
		&bigcheckbox,
		&bigcheckbox2,
		&backButton,
		&basicButton,
	}
}

func (m *SettingsScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})
	m.backButton.Draw(screen)
	m.bigcheckbox.Draw(screen)
	m.bigcheckbox2.Draw(screen)
	m.biglabel.Draw(screen)
	m.bigoutline.Draw(screen)
	m.basicButton.Draw(screen)
}

func (m *SettingsScene) Update(g *SceneManager) error {
	m.backButton.Input()
	m.biglabel.Input(g)
	m.bigcheckbox.Input()
	m.bigcheckbox2.Input()
	m.basicButton.Input()

	if m.backButton.Execute {
		g.ChangeScene("MenuScene")
		m.backButton.Execute = false
	}
	return nil
}

func (m *SettingsScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}

func (m *SettingsScene) Name() string {
	return "SettingsScene"
}
