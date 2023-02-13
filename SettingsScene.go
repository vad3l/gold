package main

import (
	"image/color"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SettingsScene struct {
	bigcontour	*Contour
	biglabel	*Label
	//bigcheckbox	*Checkbox
	backButton	*Button
}

func NewSettingsScene () *SettingsScene {


	/*bigcheckbox := Checkbox{
		0,
		[]string{ "niggure", "nighga" },
		[]Point{ Point{200, 200}, Point{300, 200} },
		Point{50, 50},
		func(g *SceneManager, i int) {
			fmt.Println("bigCheckBox")
		},
	}*/

	bigcontour := Contour{
		Point{500,500},
		Point{100,200},
		"bigContour",
	}

	biglabel := Label{
		"big label",
		Point{ 400, 10 },
		func(g *SceneManager) {
			fmt.Println("label")
		},
	}


	backButton := NewButton(Point{150,50},Point{50,50},"BACK",
	func(g *SceneManager) {
		g.current_scene = NewMenuScene()
	})
	backButton.setRadius(150)
	backButton.setColor(color.RGBA{ 0xb5, 0xf1, 0xcc, 0xff })
	backButton.setColorHover(color.RGBA{ 0xe5, 0xfd, 0xd1, 0xff })
	backButton.setColorText(color.RGBA{ 0xFF ,0xAA ,0xCF, 0xff })
	backButton.setColorTextHover(color.RGBA{ 0xFF ,0xAA ,0xCF ,0xff })
	backButton.setFont("TTF/dogicapixel.ttf")
	backButton.setFontSize(20)
	return &SettingsScene{
		&bigcontour,
		&biglabel,
		//&bigcheckbox,
		&backButton,
	}
}

func (m *SettingsScene) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0xb9, 0xf3, 0xff, 0xff })
	m.backButton.Draw(screen)
	//m.bigcheckbox.Draw(screen)
	m.biglabel.Draw(screen)
	m.bigcontour.Draw(screen)
}

func (m *SettingsScene) Update(g *SceneManager) error {
	m.backButton.Input(g)
	//m.bigcheckbox.Input(g)
	return nil
}

func (m *SettingsScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	return 1280, 720
}
