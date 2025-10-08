package gui

import (
	"image/color"
	"io/ioutil"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	. "github.com/vad3l/gold/library/graphics"
	"golang.org/x/image/font"
)

// cache pour accélérer la résolution des touches selon le layout
var charKeyCache = map[rune]ebiten.Key{}

const KeyUnknown ebiten.Key = -1

// keyForRune renvoie la touche physique correspondant au caractère 'r'
// selon la disposition du clavier (utile pour AZERTY vs QWERTY)
func keyForRune(r rune) ebiten.Key {
	if k, ok := charKeyCache[r]; ok {
		return k
	}
	for k := ebiten.KeyA; k <= ebiten.KeyZ; k++ {
		name := ebiten.KeyName(k)
		if name != "" && strings.EqualFold(name, string(r)) {
			charKeyCache[r] = k
			return k
		}
	}
	charKeyCache[r] = KeyUnknown
	return KeyUnknown
}

type TextField struct {
	Position         Point
	Size             Point
	Text             string
	Placeholder      string
	Font             font.Face
	fontParsed       *truetype.Font
	FontSize         float64
	ColorText        color.RGBA
	ColorBG          color.RGBA
	ColorBorder      color.RGBA
	ColorSelection   color.RGBA
	Active           bool
	CursorPos        int
	SelectionStart   int
	SelectionEnd     int
	MaxLength        int
	Password         bool
	NumericOnly      bool
	ScrollOffsetX    float64
	OnChange         func(*TextField)
	lastRepeat       time.Time
	repeatDelay      time.Duration
	repeatInterval   time.Duration
	cursorVisible    bool
	cursorBlinkTimer time.Time
}

func NewTextField(position, size Point) *TextField {
	return &TextField{
		Position:         position,
		Size:             size,
		Text:             "",
		Font:             nil,
		FontSize:         20,
		ColorText:        color.RGBA{0, 0, 0, 255},
		ColorBG:          color.RGBA{255, 255, 255, 255},
		ColorBorder:      color.RGBA{0, 0, 0, 255},
		ColorSelection:   color.RGBA{100, 150, 255, 128},
		MaxLength:        100,
		SelectionStart:   -1,
		SelectionEnd:     -1,
		repeatDelay:      400 * time.Millisecond,
		repeatInterval:   50 * time.Millisecond,
		cursorVisible:    true,
		cursorBlinkTimer: time.Now(),
	}
}

func (t *TextField) hasSelection() bool {
	return t.SelectionStart != -1 && t.SelectionEnd != -1 && t.SelectionStart != t.SelectionEnd
}

func (t *TextField) clearSelection() {
	t.SelectionStart = -1
	t.SelectionEnd = -1
}

func (t *TextField) deleteSelection() {
	if !t.hasSelection() {
		return
	}
	start := t.SelectionStart
	end := t.SelectionEnd
	if start > end {
		start, end = end, start
	}
	runes := []rune(t.Text)
	t.Text = string(append(runes[:start], runes[end:]...))
	t.CursorPos = start
	t.clearSelection()
	if t.OnChange != nil {
		t.OnChange(t)
	}
}

func (t *TextField) selectAll() {
	t.SelectionStart = 0
	t.SelectionEnd = utf8.RuneCountInString(t.Text)
	t.CursorPos = t.SelectionEnd
}

func (t *TextField) getDisplayText() string {
	if t.Password {
		return strings.Repeat("*", utf8.RuneCountInString(t.Text))
	}
	return t.Text
}

func (t *TextField) measureText(str string) int {
	if str == "" {
		return 0
	}
	visibleText := str
	runes := []rune(str)
	for i, r := range runes {
		if r == ' ' {
			runes[i] = '·'
		}
	}
	visibleText = string(runes)
	return text.BoundString(t.Font, visibleText).Dx()
}

func (t *TextField) Draw(screen *ebiten.Image) {
	bg := ebiten.NewImage(int(t.Size.X), int(t.Size.Y))
	bg.Fill(t.ColorBG)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.Position.X, t.Position.Y)
	screen.DrawImage(bg, op)

	displayText := t.getDisplayText()
	if displayText == "" && !t.Active && t.Placeholder != "" {
		displayText = t.Placeholder
	}

	ascent := t.Font.Metrics().Ascent.Round()
	textY := int(t.Position.Y) + ascent + 4

	opText := &ebiten.DrawImageOptions{}
	textImage := ebiten.NewImage(int(t.Size.X), int(t.Size.Y))
	textImage.Fill(color.RGBA{0, 0, 0, 0})

	if t.hasSelection() {
		start := t.SelectionStart
		end := t.SelectionEnd
		if start > end {
			start, end = end, start
		}
		beforeSel := string([]rune(displayText)[:start])
		selText := string([]rune(displayText)[start:end])

		selStartX := t.measureText(beforeSel)
		selWidth := t.measureText(beforeSel+selText) - selStartX
		selX := float64(selStartX) - t.ScrollOffsetX
		if selX < 0 {
			selWidth += int(selX)
			selX = 0
		}
		if selX+float64(selWidth) > t.Size.X {
			selWidth = int(t.Size.X - selX)
		}
		if selWidth > 0 && selX < t.Size.X {
			ebitenutil.DrawRect(textImage, selX, 0, float64(selWidth), t.Size.Y, t.ColorSelection)
		}
	}

	text.Draw(textImage, displayText, t.Font, -int(t.ScrollOffsetX), textY-int(t.Position.Y), t.ColorText)
	opText.GeoM.Translate(t.Position.X, t.Position.Y)
	screen.DrawImage(textImage, opText)

	if t.Active && !t.hasSelection() {
		if time.Since(t.cursorBlinkTimer) > 500*time.Millisecond {
			t.cursorVisible = !t.cursorVisible
			t.cursorBlinkTimer = time.Now()
		}
		if t.cursorVisible {
			cursorText := string([]rune(displayText)[:t.CursorPos])
			cursorX := int(t.Position.X) + t.measureText(cursorText) - int(t.ScrollOffsetX)
			if cursorX >= int(t.Position.X) && cursorX <= int(t.Position.X+t.Size.X) {
				ebitenutil.DrawRect(screen, float64(cursorX), t.Position.Y+2, 2, t.Size.Y-4, t.ColorText)
			}
		}
	}

	// bordure
	ebitenutil.DrawRect(screen, t.Position.X, t.Position.Y, t.Size.X, 2, t.ColorBorder)
	ebitenutil.DrawRect(screen, t.Position.X, t.Position.Y+t.Size.Y-2, t.Size.X, 2, t.ColorBorder)
	ebitenutil.DrawRect(screen, t.Position.X, t.Position.Y, 2, t.Size.Y, t.ColorBorder)
	ebitenutil.DrawRect(screen, t.Position.X+t.Size.X-2, t.Position.Y, 2, t.Size.Y, t.ColorBorder)
}

func (t *TextField) Input() {
	x, y := ebiten.CursorPosition()
	pCursor := Point{float64(x), float64(y)}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if pCursor.X >= t.Position.X && pCursor.X <= t.Position.X+t.Size.X &&
			pCursor.Y >= t.Position.Y && pCursor.Y <= t.Position.Y+t.Size.Y {
			t.Active = true
			t.CursorPos = utf8.RuneCountInString(t.Text)
			t.clearSelection()
		} else {
			t.Active = false
		}
	}

	if !t.Active {
		return
	}

	ctrlPressed := ebiten.IsKeyPressed(ebiten.KeyControl) ||
		ebiten.IsKeyPressed(ebiten.KeyControlLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyControlRight) ||
		ebiten.IsKeyPressed(ebiten.KeyMeta) ||
		ebiten.IsKeyPressed(ebiten.KeyMetaLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyMetaRight)

	shiftPressed := ebiten.IsKeyPressed(ebiten.KeyShift) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftRight)

	// ---- Gestion Ctrl/Cmd + A ----
	aKey := keyForRune('a')
	if ctrlPressed {
		if aKey != KeyUnknown {
			if inpututil.IsKeyJustPressed(aKey) {
				t.selectAll()
				return
			}
		} else {
			// fallback (rare, si KeyName ne marche pas)
			if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
				t.selectAll()
				return
			}
		}
	}

	// ---- Insertion de caractères ----
	for _, r := range ebiten.InputChars() {
		if t.hasSelection() {
			t.deleteSelection()
		}
		if t.NumericOnly && (r < '0' || r > '9') {
			continue
		}
		if utf8.RuneCountInString(t.Text) < t.MaxLength {
			runes := []rune(t.Text)
			runes = append(runes[:t.CursorPos], append([]rune{r}, runes[t.CursorPos:]...)...)
			t.Text = string(runes)
			t.CursorPos++
			if t.OnChange != nil {
				t.OnChange(t)
			}
			t.adjustScroll()
		}
	}

	// ---- Navigation / suppression ----
	now := time.Now()
	keys := []struct {
		key ebiten.Key
		fn  func()
	}{
		{ebiten.KeyBackspace, func() {
			if t.hasSelection() {
				t.deleteSelection()
				t.adjustScroll()
			} else if t.CursorPos > 0 {
				runes := []rune(t.Text)
				t.Text = string(append(runes[:t.CursorPos-1], runes[t.CursorPos:]...))
				t.CursorPos--
				if t.OnChange != nil {
					t.OnChange(t)
				}
				t.adjustScroll()
			}
		}},
		{ebiten.KeyDelete, func() {
			if t.hasSelection() {
				t.deleteSelection()
				t.adjustScroll()
			} else if t.CursorPos < utf8.RuneCountInString(t.Text) {
				runes := []rune(t.Text)
				t.Text = string(append(runes[:t.CursorPos], runes[t.CursorPos+1:]...))
				if t.OnChange != nil {
					t.OnChange(t)
				}
				t.adjustScroll()
			}
		}},
		{ebiten.KeyArrowLeft, func() {
			if t.CursorPos > 0 {
				if shiftPressed {
					if !t.hasSelection() {
						t.SelectionStart = t.CursorPos
					}
					t.CursorPos--
					t.SelectionEnd = t.CursorPos
				} else {
					if t.hasSelection() {
						start := t.SelectionStart
						end := t.SelectionEnd
						if start > end {
							start, end = end, start
						}
						t.CursorPos = start
						t.clearSelection()
					} else {
						t.CursorPos--
					}
				}
				t.adjustScroll()
			}
		}},
		{ebiten.KeyArrowRight, func() {
			if t.CursorPos < utf8.RuneCountInString(t.Text) {
				if shiftPressed {
					if !t.hasSelection() {
						t.SelectionStart = t.CursorPos
					}
					t.CursorPos++
					t.SelectionEnd = t.CursorPos
				} else {
					if t.hasSelection() {
						start := t.SelectionStart
						end := t.SelectionEnd
						if start > end {
							start, end = end, start
						}
						t.CursorPos = end
						t.clearSelection()
					} else {
						t.CursorPos++
					}
				}
				t.adjustScroll()
			}
		}},
	}

	for _, k := range keys {
		if ebiten.IsKeyPressed(k.key) {
			if now.Sub(t.lastRepeat) > t.repeatDelay {
				k.fn()
				t.lastRepeat = now
			} else if now.Sub(t.lastRepeat) > t.repeatInterval {
				k.fn()
				t.lastRepeat = now
			}
		}
	}
}

func (t *TextField) adjustScroll() {
	displayText := t.getDisplayText()
	cursorText := string([]rune(displayText)[:t.CursorPos])

	cursorPixel := t.measureText(cursorText)
	width := t.measureText(displayText)

	if float64(cursorPixel)-t.ScrollOffsetX > t.Size.X-10 {
		t.ScrollOffsetX = float64(cursorPixel) - t.Size.X + 10
	} else if float64(cursorPixel)-t.ScrollOffsetX < 0 {
		t.ScrollOffsetX = float64(cursorPixel)
	}
	if float64(width) <= t.Size.X {
		t.ScrollOffsetX = 0
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
	t.ScrollOffsetX = 0
	t.clearSelection()
}
