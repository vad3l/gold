package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/vad3l/gold/library/graphics"
	. "github.com/vad3l/gold/library/graphics/gui"
)

type SettingsScene struct {
	bigOutline   *Outline
	bigLabel     *Label
	bigCheckbox  *CheckBox
	bigCheckbox2 *CheckBox
	backButton   *Button
	basicButton  *Button
	textField    *TextField
	listView     *ListView
}

func NewSettingsScene() *SettingsScene {

	bigCheckbox := NewCheckBox(Point{30, 30}, Point{100, 400})
	bigCheckbox.SetRadius(8)
	bigCheckbox.SetBorderSize(5)
	bigCheckbox.SetColor(color.RGBA{0, 0, 200, 255})
	bigCheckbox.SetColorChecked(color.RGBA{0, 200, 0, 255})

	bigCheckbox2 := NewCheckBox(Point{30, 30}, Point{150, 450})
	bigCheckbox2.SetRadius(8)
	bigCheckbox2.SetBorderSize(5)
	bigCheckbox2.SetColor(color.RGBA{0, 0, 200, 255})
	bigCheckbox2.SetColorChecked(color.RGBA{0, 200, 0, 255})

	radioGroup := NewRadioGroup()
	bigCheckbox.SetRadioGroup(radioGroup)
	bigCheckbox2.SetRadioGroup(radioGroup)
	radioGroup.Add(&bigCheckbox)
	radioGroup.Add(&bigCheckbox2)

	bigOutline := Outline{
		Size:     Point{500, 500},
		Position: Point{100, 200},
		Text:     "bigoutline",
	}

	bigLabel := NewLabel("To Show example push 1 2 3", Point{0, 0})
	bigLabel.SetFunction(func() {
		fmt.Println("label clicked")
	})
	bigLabel.SetColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	bigLabel.SetFont("data/dogicapixel.ttf")
	bigLabel.SetFontSize(50)

	backButton := NewButton(Point{150, 50}, Point{50, 50}, "BACK")
	backButton.SetRadius(150)
	backButton.SetColor(color.RGBA{0xb5, 0xf1, 0xcc, 0xff})
	backButton.SetColorHover(color.RGBA{0xe5, 0xfd, 0xd1, 0xff})
	backButton.SetColorText(color.RGBA{0xFF, 0xAA, 0xCF, 0xff})
	backButton.SetColorTextHover(color.RGBA{0xFF, 0xAA, 0xCF, 0xff})
	backButton.SetFont("data/dogicapixel.ttf")
	backButton.SetFontSize(20)

	basicButton := NewButton(Point{100, 200}, Point{800, 500}, "Basic")

	textField := NewTextField(Point{300, 600}, Point{400, 40})
	textField.SetBackgroundColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	textField.SetTextColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	textField.SetPlaceholder("Enter text here...")
	textField.SetFont("data/dogicapixel.ttf")
	textField.SetFontSize(20)

	// ListView
	listView := NewListView(Point{600, 200}, Point{200, 300})
	for i := 1; i <= 10; i++ {
		btn := NewButton(Point{0, 0}, Point{180, 40}, fmt.Sprintf("Item %d", i))
		btn.SetFont("data/dogicapixel.ttf")
		btn.SetFontSize(16)
		listView.Add(&btn)
	}

	return &SettingsScene{
		bigOutline:   &bigOutline,
		bigLabel:     &bigLabel,
		bigCheckbox:  &bigCheckbox,
		bigCheckbox2: &bigCheckbox2,
		backButton:   &backButton,
		basicButton:  &basicButton,
		textField:    textField,
		listView:     listView,
	}
}

func (m *SettingsScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})

	m.backButton.Draw(screen)
	m.bigCheckbox.Draw(screen)
	m.bigCheckbox2.Draw(screen)
	m.bigLabel.Draw(screen)
	m.bigOutline.Draw(screen)
	m.basicButton.Draw(screen)
	m.textField.Draw(screen)
	m.listView.Draw(screen)
}

func (m *SettingsScene) Update(g *SceneManager) error {
	m.backButton.Input()
	m.bigLabel.Input()
	m.bigCheckbox.Input()
	m.bigCheckbox2.Input()
	m.basicButton.Input()
	m.textField.Input()
	m.listView.Input()

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
