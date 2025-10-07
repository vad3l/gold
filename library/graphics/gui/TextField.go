package gui

import (
	"image/color"
	"io/ioutil"
	"strings"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	. "github.com/vad3l/gold/library/graphics"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type TextField struct {
	Position       Point
	Size           Point
	Text           string
	Placeholder    string
	Font           font.Face
	fontParsed     *truetype.Font
	FontSize       float64
	ColorText      color.RGBA
	ColorBG        color.RGBA
	ColorBorder    color.RGBA
	Active         bool
	CursorPos      int
	SelectionStart int
	SelectionEnd   int
	MaxLength      int
	Password       bool
	NumericOnly    bool
	Multiline      bool
	Mask           func(r rune) rune
	OnChange       func(*TextField)
	ScrollOffsetX  float64
}

func NewTextField(position, size Point) *TextField {
	return &TextField{
		Position:    position,
		Size:        size,
		Text:        "",
		Font:        basicfont.Face7x13,
		FontSize:    12,
		ColorText:   color.RGBA{0, 0, 0, 255},
		ColorBG:     color.RGBA{255, 255, 255, 255},
		ColorBorder: color.RGBA{0, 0, 0, 255},
		MaxLength:   100,
	}
}

func (t *TextField) Draw(screen *ebiten.Image) {
	bg := ebiten.NewImage(int(t.Size.X), int(t.Size.Y))
	bg.Fill(t.ColorBG)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.Position.X, t.Position.Y)
	screen.DrawImage(bg, op)

	displayText := t.Text
	if t.Password {
		displayText = strings.Repeat("*", utf8.RuneCountInString(t.Text))
	}
	if displayText == "" && !t.Active && t.Placeholder != "" {
		displayText = t.Placeholder
	}

	textX := t.Position.X - t.ScrollOffsetX
	textY := t.Position.Y + int(t.Size.Y/2)
	text.Draw(screen, displayText, t.Font, int(textX), int(textY), t.ColorText)

	if t.Active {
		cursorX := t.Position.X - t.ScrollOffsetX + float64(text.BoundString(t.Font, string([]rune(displayText)[:t.CursorPos])).Dx())
		ebitenutil.DrawRect(screen, cursorX, t.Position.Y+2, 2, t.Size.Y-4, color.RGBA{0, 0, 0, 255})
	}
	ebitenutil.DrawRect(screen, t.Position.X, t.Position.Y, t.Size.X, t.Size.Y, t.ColorBorder)
}

func (t *TextField) Input() {
	x, y := ebiten.CursorPosition()
	pCursor := Point{float64(x), float64(y)}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if pCursor.X >= t.Position.X && pCursor.X <= t.Position.X+t.Size.X &&
			pCursor.Y >= t.Position.Y && pCursor.Y <= t.Position.Y+t.Size.Y {
			t.Active = true
			t.CursorPos = utf8.RuneCountInString(t.Text)
		} else {
			t.Active = false
		}
	}

	if t.Active {
		for _, r := range inpututil.InputChars() {
			if t.NumericOnly && (r < '0' || r > '9') {
				continue
			}
			if t.Mask != nil {
				r = t.Mask(r)
			}
			if utf8.RuneCountInString(t.Text) < t.MaxLength {
				runes := []rune(t.Text)
				if t.SelectionStart != t.SelectionEnd {
					start := min(t.SelectionStart, t.SelectionEnd)
					end := max(t.SelectionStart, t.SelectionEnd)
					runes = append(runes[:start], runes[end:]...)
					t.SelectionStart = start
					t.SelectionEnd = start
					t.CursorPos = start
				}
				runes = append(runes[:t.CursorPos], append([]rune{r}, runes[t.CursorPos:]...)...)
				t.Text = string(runes)
				t.CursorPos++
				if t.OnChange != nil {
					t.OnChange(t)
				}
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if t.SelectionStart != t.SelectionEnd {
				start := min(t.SelectionStart, t.SelectionEnd)
				end := max(t.SelectionStart, t.SelectionEnd)
				runes := []rune(t.Text)
				runes = append(runes[:start], runes[end:]...)
				t.Text = string(runes)
				t.CursorPos = start
				t.SelectionStart = start
				t.SelectionEnd = start
			} else if t.CursorPos > 0 {
				runes := []rune(t.Text)
				t.Text = string(append(runes[:t.CursorPos-1], runes[t.CursorPos:]...))
				t.CursorPos--
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyDelete) && t.CursorPos < utf8.RuneCountInString(t.Text) {
			runes := []rune(t.Text)
			t.Text = string(append(runes[:t.CursorPos], runes[t.CursorPos+1:]...))
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && t.CursorPos > 0 {
			t.CursorPos--
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && t.CursorPos < utf8.RuneCountInString(t.Text) {
			t.CursorPos++
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
	t.Font = truetype.NewFace(t.fontParsed, &truetype.Options{
		Size:    t.FontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func (t *TextField) SetFontSize(size float64) {
	t.FontSize = size
	if t.fontParsed != nil {
		t.Font = truetype.NewFace(t.fontParsed, &truetype.Options{
			Size:    t.FontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
	}
}

func (t *TextField) SetTextColor(c color.RGBA) {
	t.ColorText = c
}

func (t *TextField) SetBackgroundColor(c color.RGBA) {
	t.ColorBG = c
}

func (t *TextField) SetBorderColor(c color.RGBA) {
	t.ColorBorder = c
}

func (t *TextField) SetPlaceholder(text string) {
	t.Placeholder = text
}

func (t *TextField) Clear() {
	t.Text = ""
	t.CursorPos = 0
	t.SelectionStart = 0
	t.SelectionEnd = 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
