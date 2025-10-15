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
	// example buttons demonstrating wrapping and truncation
	wrapButton  *Button
	truncButton *Button
}

func NewSettingsScene() *SettingsScene {
	// Main panel outline (centered-ish)
	panelPos := Point{X: 80, Y: 80}
	panelSize := Point{X: 1120, Y: 560}
	bigOutline := Outline{
		Size:     panelSize,
		Position: panelPos,
		Text:     "",
	}

	// Title label - draw near top center of the panel
	bigLabel := NewLabel("Settings", Point{X: panelPos.X + panelSize.X/2 - 40, Y: panelPos.Y + 18})
	bigLabel.SetFunction(func() {
		fmt.Println("label clicked")
	})
	bigLabel.SetColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	bigLabel.SetFont("data/dogicapixel.ttf")
	bigLabel.SetFontSize(28)

	// Left column: radio options with labels
	leftX := panelPos.X + 30
	leftY := panelPos.Y + 70
	radioGroup := NewRadioGroup()

	opt1 := NewCheckBox(Point{X: 20, Y: 20}, Point{X: leftX, Y: leftY})
	opt1.SetRadius(4)
	opt1.SetBorderSize(2)
	opt1.SetColor(color.RGBA{0x88, 0x88, 0xff, 0xff})
	opt1.SetColorChecked(color.RGBA{0x44, 0xaa, 0x44, 0xff})
	opt1.SetRadioGroup(radioGroup)
	radioGroup.Add(&opt1)

	label1 := NewLabel("Option A (radio)", Point{X: leftX + 28, Y: leftY - 4})
	label1.SetFont("data/dogicapixel.ttf")
	label1.SetFontSize(16)

	opt2 := NewCheckBox(Point{X: 20, Y: 20}, Point{X: leftX, Y: leftY + 40})
	opt2.SetRadius(4)
	opt2.SetBorderSize(2)
	opt2.SetColor(color.RGBA{0x88, 0x88, 0xff, 0xff})
	opt2.SetColorChecked(color.RGBA{0x44, 0xaa, 0x44, 0xff})
	opt2.SetRadioGroup(radioGroup)
	radioGroup.Add(&opt2)

	label2 := NewLabel("Option B (radio)", Point{X: leftX + 28, Y: leftY + 36})
	label2.SetFont("data/dogicapixel.ttf")
	label2.SetFontSize(16)

	// Small description label under radios
	desc := NewLabel("Choose one of the options on the left.", Point{X: leftX, Y: leftY + 84})
	desc.SetFont("data/dogicapixel.ttf")
	desc.SetFontSize(14)

	// Horizontal text field under left column
	tf := NewTextField(Point{X: leftX, Y: leftY + 120}, Point{X: 360, Y: 28})
	tf.SetBackgroundColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	tf.SetTextColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	tf.SetPlaceholder("Enter value...")
	tf.SetFont("data/dogicapixel.ttf")
	tf.SetFontSize(14)

	// Bottom-left action buttons
	backButton := NewButton(Point{X: 100, Y: 36}, Point{X: leftX, Y: panelPos.Y + panelSize.Y - 60}, "Back")
	backButton.SetRadius(6)
	backButton.SetColor(color.RGBA{0xd3, 0xe8, 0xff, 0xff})
	backButton.SetColorHover(color.RGBA{0xbf, 0xdf, 0xff, 0xff})
	backButton.SetFont("data/dogicapixel.ttf")
	backButton.SetFontSize(16)

	applyButton := NewButton(Point{X: 100, Y: 36}, Point{X: leftX + 120, Y: panelPos.Y + panelSize.Y - 60}, "Apply")
	applyButton.SetRadius(6)
	applyButton.SetColor(color.RGBA{0xc8, 0xff, 0xd0, 0xff})
	applyButton.SetColorHover(color.RGBA{0xaa, 0xff, 0xc2, 0xff})
	applyButton.SetFont("data/dogicapixel.ttf")
	applyButton.SetFontSize(16)

	// Right panel ListView (smaller, navigator-style)
	rvX := panelPos.X + panelSize.X - 360
	rvY := panelPos.Y + 60
	listView := NewListView(Point{X: rvX, Y: rvY}, Point{X: 320, Y: panelSize.Y - 120})
	for i := 1; i <= 8; i++ {
		btn := NewButton(Point{X: 300, Y: 36}, Point{X: 0, Y: 0}, fmt.Sprintf("Setting %d", i))
		btn.SetFont("data/dogicapixel.ttf")
		btn.SetFontSize(14)
		listView.Add(&btn)
	}

	// Example buttons demonstrating wrapping and truncation
	wrapBtn := NewButton(Point{X: 360, Y: 40}, Point{X: leftX, Y: leftY + 160}, "This is a very long button label that will wrap to multiple lines when WrapText is enabled")
	wrapBtn.SetFont("data/dogicapixel.ttf")
	wrapBtn.SetFontSize(14)
	wrapBtn.WrapText = true
	wrapBtn.Ellipsis = false
	wrapBtn.Padding = 8

	truncBtn := NewButton(Point{X: 360, Y: 40}, Point{X: leftX, Y: leftY + 220}, "This is another very long label that will be truncated if it doesn't fit the button width")
	truncBtn.SetFont("data/dogicapixel.ttf")
	truncBtn.SetFontSize(14)
	truncBtn.WrapText = false
	truncBtn.Ellipsis = true
	truncBtn.Padding = 8

	return &SettingsScene{
		bigOutline:   &bigOutline,
		bigLabel:     &bigLabel,
		bigCheckbox:  &opt1,
		bigCheckbox2: &opt2,
		backButton:   &backButton,
		basicButton:  &applyButton,
		textField:    tf,
		listView:     listView,
		wrapButton:   &wrapBtn,
		truncButton:  &truncBtn,
	}
}

func (m *SettingsScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xb9, 0xf3, 0xff, 0xff})

	// Draw main panel
	m.bigOutline.Draw(screen)

	// Header and back
	m.backButton.Draw(screen)
	m.bigLabel.Draw(screen)

	// Left column controls (inside panel)
	m.bigCheckbox.Draw(screen)
	m.bigCheckbox2.Draw(screen)
	m.textField.Draw(screen)
	m.basicButton.Draw(screen)
	// example buttons
	if m.wrapButton != nil {
		m.wrapButton.Draw(screen)
	}
	if m.truncButton != nil {
		m.truncButton.Draw(screen)
	}

	// Right panel list
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
	if m.wrapButton != nil {
		m.wrapButton.Input()
	}
	if m.truncButton != nil {
		m.truncButton.Input()
	}

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
