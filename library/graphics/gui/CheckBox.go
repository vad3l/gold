package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	. "github.com/vad3l/gold/library/graphics"
)

type CheckBox struct {
	Size Point

	position             Point
	colorCheckBox        color.RGBA
	colorCheckBoxChecked color.RGBA
	borderSize           float64
	Checked              bool
	Execute              bool
	radius               int

	img        *ebiten.Image
	radioGroup *RadioGroup // Ajout pour le mode radio
}

func NewCheckBox(size, position Point) CheckBox {

	return CheckBox{
		size,
		position,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		10.0,
		false,
		false,
		0,
		nil,
		nil,
	}
}

func (c *CheckBox) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(c.Size.X), int(c.Size.Y))
	c.img = img

	if c.radius == 0 {
		ebitenutil.DrawRect(img, 0, 0, c.Size.X, c.Size.Y, c.colorCheckBox)
		if c.Checked {
			ebitenutil.DrawRect(img, c.borderSize, c.borderSize, c.Size.X-c.borderSize*2, c.Size.Y-c.borderSize*2, c.colorCheckBoxChecked)
		}
	} else {
		ebitenutil.DrawRect(img, 0, 0+float64(c.radius), c.Size.X, c.Size.Y-2*float64(c.radius), c.colorCheckBox)
		ebitenutil.DrawRect(img, 0+float64(c.radius), 0, c.Size.X-2*float64(c.radius), c.Size.Y, c.colorCheckBox)

		ebitenutil.DrawCircle(img, 0+float64(c.radius), 0+float64(c.radius), float64(c.radius), c.colorCheckBox)
		ebitenutil.DrawCircle(img, 0+c.Size.X-float64(c.radius), 0+float64(c.radius), float64(c.radius), c.colorCheckBox)
		ebitenutil.DrawCircle(img, 0+float64(c.radius), 0+c.Size.Y-float64(c.radius), float64(c.radius), c.colorCheckBox)
		ebitenutil.DrawCircle(img, 0+c.Size.X-float64(c.radius), 0+c.Size.Y-float64(c.radius), float64(c.radius), c.colorCheckBox)
		if c.Checked {
			ebitenutil.DrawRect(img, c.borderSize, c.borderSize+float64(c.radius), c.Size.X-c.borderSize*2, c.Size.Y-c.borderSize*2-2*float64(c.radius), c.colorCheckBoxChecked)
			ebitenutil.DrawRect(img, c.borderSize+float64(c.radius), c.borderSize, c.Size.X-c.borderSize*2-2*float64(c.radius), c.Size.Y-c.borderSize*2, c.colorCheckBoxChecked)

			ebitenutil.DrawCircle(img, c.borderSize+float64(c.radius), c.borderSize+float64(c.radius), float64(c.radius), c.colorCheckBoxChecked)
			ebitenutil.DrawCircle(img, c.borderSize+c.Size.X-c.borderSize*2-float64(c.radius), c.borderSize+float64(c.radius), float64(c.radius), c.colorCheckBoxChecked)
			ebitenutil.DrawCircle(img, c.borderSize+float64(c.radius), c.borderSize+c.Size.Y-c.borderSize*2-float64(c.radius), float64(c.radius), c.colorCheckBoxChecked)
			ebitenutil.DrawCircle(img, c.borderSize+c.Size.X-c.borderSize*2-float64(c.radius), c.borderSize+c.Size.Y-c.borderSize*2-float64(c.radius), float64(c.radius), c.colorCheckBoxChecked)
		}
	}

	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Translate(c.position.X, c.position.Y)
	c.img = img
	screen.DrawImage(img, ot)
}

func (c *CheckBox) Input() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pCursor := Point{float64(x), float64(y)}
		if hover(pCursor, c.Size, c.position, c.img, 1) {
			if c.radioGroup != nil {
				c.radioGroup.Select(c)
			} else {
				c.Checked = !c.Checked
				c.Execute = true
			}
		}
	}
}

func (c *CheckBox) SetRadioGroup(group *RadioGroup) {
	c.radioGroup = group
}

func (c *CheckBox) SetRadius(radius int) {
	if radius > int(c.Size.X/2)-int(c.borderSize) {
		c.radius = int(c.Size.X/2) - int(c.borderSize)
		return
	}
	if radius > int(c.Size.Y/2) {
		c.radius = int(c.Size.Y/2) - int(c.borderSize)
		return
	} else {
		c.radius = radius
	}
}

func (c *CheckBox) SetColor(color color.RGBA) {
	c.colorCheckBox = color
}

func (c *CheckBox) SetColorChecked(color color.RGBA) {
	c.colorCheckBoxChecked = color
}

func (c *CheckBox) SetBorderSize(border float64) {
	if border < 0 {
		c.borderSize = 0
		return
	}
	c.borderSize = border
}

type RadioGroup struct {
	checkBoxes []*CheckBox
}

func NewRadioGroup() *RadioGroup {
	return &RadioGroup{
		checkBoxes: []*CheckBox{},
	}
}

func (g *RadioGroup) Add(cb *CheckBox) {
	g.checkBoxes = append(g.checkBoxes, cb)
}

func (g *RadioGroup) Select(selected *CheckBox) {
	for _, cb := range g.checkBoxes {
		if cb == selected {
			cb.Checked = true
			cb.Execute = true
		} else {
			cb.Checked = false
			cb.Execute = false
		}
	}
}
