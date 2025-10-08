package gui

import (
	"fmt"
	"image/color"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"

	. "github.com/vad3l/gold/library/graphics"
)

type Label struct {
	text     string
	Position Point
	execute  func()

	font           font.Face
	fontSize       float64
	fontParsed     *truetype.Font
	colorText      color.RGBA
	colorTextHover color.RGBA
}

func NewLabel(text string, position Point) Label {
	return Label{
		text:           text,
		Position:       position,
		execute:        nil,
		font:           basicfont.Face7x13,
		fontSize:       12,
		fontParsed:     nil,
		colorText:      color.RGBA{0x00, 0x00, 0x00, 0xff},
		colorTextHover: color.RGBA{0x00, 0x00, 0x00, 0xff},
	}
}

func (l *Label) Draw(screen *ebiten.Image) {
	fontDimension := text.BoundString(l.font, l.text)
	sum := (fontDimension.Max.Y + fontDimension.Max.X) / 2
	text.Draw(screen, l.text, l.font, int(l.Position.X), int(l.Position.Y)+sum/4, l.colorText)
}

func (l *Label) Input() {
	if l.execute != nil && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		size := Point{float64(7 * len(l.text)), 13}
		if float64(x) >= l.Position.X && float64(x) <= (l.Position.X+size.X) &&
			float64(y) >= l.Position.Y && float64(y) <= (l.Position.Y+size.Y) {
			l.execute()
		}
	}
}

func (l *Label) SetFunction(execute func()) {
	l.execute = execute
}

func (l *Label) SetText(text string) {
	l.text = text
}

func (l *Label) SetFont(path string) {
	tt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fontParsed, erre := truetype.Parse(tt)
	if erre != nil {
		panic(erre)
	}
	l.fontParsed = fontParsed

	fontFace := truetype.NewFace(l.fontParsed, &truetype.Options{
		Size:    l.fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	l.font = fontFace
}

func (l *Label) SetFontSize(size float64) {
	l.fontSize = size
	if l.fontParsed == nil {
		fmt.Println("Warning - as long as you haven't changed the font, it can't change its size")
	} else {
		fontFace := truetype.NewFace(l.fontParsed, &truetype.Options{
			Size:    l.fontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		l.font = fontFace
	}
}

func (l *Label) SetColor(colorButton color.RGBA) {
	l.colorText = colorButton
}

func (l *Label) SetPosition(position Point) {
	l.Position = position
}
