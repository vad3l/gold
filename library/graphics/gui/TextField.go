package gui

import (
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"

	. "github.com/vad3l/gold/library/graphics"
)

type TextField struct {
	Position    Point
	Size        Point
	Text        string
	fontParsed  *truetype.Font
	Font        font.Face
	FontSize    float64
	ColorText   color.RGBA
	ColorBG     color.RGBA
	Active      bool
	MaxLength   int
	placeholder string
}

func NewTextField(position, size Point) *TextField {
	return &TextField{
		Position:  position,
		Size:      size,
		Text:      "",
		Font:      basicfont.Face7x13,
		FontSize:  12,
		ColorText: color.RGBA{0, 0, 0, 255},
		ColorBG:   color.RGBA{255, 255, 255, 255},
		Active:    false,
		MaxLength: 100,
	}
}

func (t *TextField) Draw(screen *ebiten.Image) {
	bg := ebiten.NewImage(int(t.Size.X), int(t.Size.Y))
	bg.Fill(t.ColorBG)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.Position.X, t.Position.Y)
	screen.DrawImage(bg, op)

	displayText := t.Text
	if t.Text == "" && !t.Active && t.placeholder != "" {
		displayText = t.placeholder
	}
	text.Draw(screen, displayText, t.Font, int(t.Position.X), int(t.Position.Y+t.Size.Y/2), t.ColorText)
}

func (t *TextField) Input() {
	x, y := ebiten.CursorPosition()
	pCursor := Point{float64(x), float64(y)}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if pCursor.X >= t.Position.X && pCursor.X <= t.Position.X+t.Size.X &&
			pCursor.Y >= t.Position.Y && pCursor.Y <= t.Position.Y+t.Size.Y {
			t.Active = true
		} else {
			t.Active = false
		}
	}

	if t.Active {
		for _, r := range inpututil.AppendPressedKeys(nil) {
			if r == 8 || r == 127 {
				if len(t.Text) > 0 {
					t.Text = t.Text[:len(t.Text)-1]
				}
			} else if len(t.Text) < t.MaxLength {
				if r >= 32 && r <= 126 {
					t.Text += string(r)
				}
			}
		}
	}
}

func (t *TextField) SetFont(path string) {
	tt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fontParsed, erre := truetype.Parse(tt)
	if erre != nil {
		panic(erre)
	}
	t.fontParsed = fontParsed

	fontFace := truetype.NewFace(t.fontParsed, &truetype.Options{
		Size:    t.FontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	t.Font = fontFace
}

func (t *TextField) SetFontSize(size float64) {
	t.FontSize = size
}

func (t *TextField) SetTextColor(c color.RGBA) {
	t.ColorText = c
}

func (t *TextField) SetBackgroundColor(c color.RGBA) {
	t.ColorBG = c
}

func (t *TextField) SetPlaceholder(text string) {
	t.placeholder = text
}

func (t *TextField) Clear() {
	t.Text = ""
}
