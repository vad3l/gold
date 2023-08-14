package gui

import (
	"image/color"
	"fmt"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type Label struct {
	text string
	Position Point
	execute func(g *SceneManager)

	font	font.Face
	fontSize	float64
	fontParsed	*truetype.Font
	colorText	color.RGBA
	colorTextHover	color.RGBA
}

func NewLabel (text string, position Point) Label {
	return Label{
		text,
		position,
		nil,
		basicfont.Face7x13,
		12,
		nil,
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0x00, 0xff},
	}
}


func (l *Label) Draw (screen *ebiten.Image) {
	fontDimension := text.BoundString(l.font,l.text)
	sum := (fontDimension.Max.Y + fontDimension.Max.X )/2
	text.Draw(screen, l.text, l.font, int(l.Position.X), int(l.Position.Y)+sum/4, l.colorText)
}

func (l *Label) Input (g *SceneManager) {
	if (l.execute != nil){
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

			x, y := ebiten.CursorPosition()
			size := Point{ float64(7 * len(l.text)), 13 }
			if float64(x) >= l.Position.X && float64(x) <= (l.Position.X + size.X) && float64(y) >= l.Position.Y && float64(y) <= (l.Position.Y + size.Y) {
				l.execute(g)
			}
		}

	}
}

func (l *Label) SetFunction (execute func(g *SceneManager)) {
	l.execute = execute
}

func (l *Label) SetText (text string) {
	l.text = text
}

func (l *Label) SetFont (path string) {
	tt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fontParsed, erre := truetype.Parse(tt)
	if erre != nil {
		panic(erre)
	}
	l.fontParsed = fontParsed

	// Create a font face with a specific size
	fontFace := truetype.NewFace(l.fontParsed, &truetype.Options{
		Size:    l.fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	l.font = fontFace
}

func (l *Label) SetFontSize (size float64) {
	l.fontSize = size
	if (l.fontParsed == nil){
		fmt.Println("Warning - as long as you haven't changed the font, it can't change its size")
	}else{
		fontFace := truetype.NewFace(l.fontParsed, &truetype.Options{
			Size:    l.fontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		l.font = fontFace
	}
	
	
}

func (l *Label) SetColor (colorButton color.RGBA) {
	l.colorText = colorButton
}