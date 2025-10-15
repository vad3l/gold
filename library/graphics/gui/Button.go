package gui

import (
	"fmt"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"

	"github.com/golang/freetype/truetype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	gfx "github.com/vad3l/gold/library/graphics"
)

type Button struct {
	Size             gfx.Point
	position         gfx.Point
	colorButton      color.RGBA
	colorButtonHover color.RGBA
	borderSize       float64
	text             string
	font             font.Face
	fontSize         float64
	fontParsed       *truetype.Font
	colorText        color.RGBA
	colorTextHover   color.RGBA
	radius           int
	Execute          bool
	img              *ebiten.Image
	// WrapText enables word-wrapping inside the button. When enabled,
	// the button will grow vertically to fit the wrapped lines.
	WrapText bool
	// Ellipsis will add "..." when truncating text (only used when WrapText==false)
	Ellipsis bool
	// Padding around the text (pixels)
	Padding int
	// extra pixels between lines when wrapped
	LineSpacing int
}

func NewButton(size, position gfx.Point, text string) Button {
	return Button{
		size,
		position,
		// base: soft pale blue to match Outline panel
		color.RGBA{0xd3, 0xe8, 0xff, 0xff},
		// hover: slightly deeper blue
		color.RGBA{0xbf, 0xdf, 0xff, 0xff},
		2.0,
		text,
		basicfont.Face7x13,
		14,
		nil,
		color.RGBA{0x10, 0x17, 0x17, 0xff},
		color.RGBA{0x06, 0x0b, 0x0b, 0xff},
		0,
		false,
		nil,
		false,
		true,
		6,
		4,
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	width := int(math.Max(b.Size.X, 1))
	// We'll compute renderHeight later (may grow if WrapText is enabled)
	baseHeight := int(math.Max(b.Size.Y, 1))
	var img *ebiten.Image
	x, y := ebiten.CursorPosition()
	pCursor := gfx.Point{X: float64(x), Y: float64(y)}
	// compute available content width (inside paddings)
	padding := b.Padding
	if padding < 0 {
		padding = 0
	}
	availWidth := width - padding*2

	// helper to measure string width using current font
	measure := func(s string) int {
		d := text.BoundString(b.font, s)
		return d.Max.X - d.Min.X
	}

	// line height using font metrics
	ascent := int(b.font.Metrics().Ascent >> 6)
	descent := int(b.font.Metrics().Descent >> 6)
	lineHeight := ascent + descent

	lines := []string{b.text}
	if b.WrapText && availWidth > 10 {
		// word-wrap into lines that fit availWidth
		words := strings.Fields(b.text)
		var cur strings.Builder
		lines = []string{}
		for i, w := range words {
			if cur.Len() == 0 {
				cur.WriteString(w)
			} else {
				// try with a space + word
				test := cur.String() + " " + w
				if measure(test) <= availWidth {
					cur.WriteString(" ")
					cur.WriteString(w)
				} else {
					// push current and start new
					lines = append(lines, cur.String())
					cur.Reset()
					cur.WriteString(w)
				}
			}
			// if last word, push remaining
			if i == len(words)-1 && cur.Len() > 0 {
				lines = append(lines, cur.String())
			}
		}
		// If no words (empty) ensure single empty line
		if len(lines) == 0 {
			lines = []string{""}
		}
	} else {
		// no wrap: possibly truncate to available width
		if availWidth > 10 {
			if measure(b.text) > availWidth {
				if b.Ellipsis {
					// leave space for ellipsis
					ell := "..."
					ellW := measure(ell)
					// binary search max prefix length that fits availWidth-ellW
					lo, hi := 0, len(b.text)
					for lo < hi {
						mid := (lo + hi + 1) / 2
						if measure(b.text[:mid]) <= (availWidth - ellW) {
							lo = mid
						} else {
							hi = mid - 1
						}
					}
					if lo < len(b.text) {
						lines = []string{b.text[:lo] + ell}
					} else {
						lines = []string{b.text}
					}
				} else {
					// truncate without ellipsis
					lo, hi := 0, len(b.text)
					for lo < hi {
						mid := (lo + hi + 1) / 2
						if measure(b.text[:mid]) <= availWidth {
							lo = mid
						} else {
							hi = mid - 1
						}
					}
					lines = []string{b.text[:lo]}
				}
			} else {
				lines = []string{b.text}
			}
		}
	}

	// compute needed height for lines + paddings (account for LineSpacing between lines)
	extraSpacing := 0
	if len(lines) > 1 {
		extraSpacing = b.LineSpacing * (len(lines) - 1)
	}
	neededHeight := lineHeight*len(lines) + extraSpacing + padding*2
	renderHeight := baseHeight
	if b.WrapText && neededHeight > baseHeight {
		renderHeight = neededHeight
		// update actual size so subsequent layout can adapt
		b.Size.Y = float64(renderHeight)
	}
	height := renderHeight
	img = ebiten.NewImage(width, height)
	b.img = img
	if b.radius == 0 {
		// drop shadow
		shadowCol := color.RGBA{0x00, 0x00, 0x00, 0x30}
		ebitenutil.DrawRect(img, 3, 3, b.Size.X, float64(height), shadowCol)
		// base
		ebitenutil.DrawRect(img, 0, 0, b.Size.X, float64(height), b.colorButton)
		// top highlight (subtle)
		highCol := b.colorButton
		highCol.A = 0x40
		ebitenutil.DrawRect(img, 2, 2, b.Size.X-4, float64(height)/3, highCol)
		if hover(pCursor, b.Size, b.position, img, 1) {
			// hovered inner fill
			ebitenutil.DrawRect(img, b.borderSize, b.borderSize, b.Size.X-b.borderSize*2, float64(height)-b.borderSize*2, b.colorButtonHover)
			// glow border
			glow := b.colorButtonHover
			if glow.A > 0x80 {
				glow.A = 0x80
			}
			ebitenutil.DrawRect(img, 0, 0, b.Size.X, float64(2), glow)
		}
		// draw lines centered horizontally and vertically (account for LineSpacing)
		totalTextHeight := lineHeight*len(lines) + extraSpacing
		// baseline start (top Y for first line's baseline)
		startY := (height-totalTextHeight)/2 + ascent
		for i, line := range lines {
			d := text.BoundString(b.font, line)
			lineW := d.Max.X - d.Min.X
			tx := int(float64(width)/2.0) - lineW/2
			ty := startY + i*(lineHeight+b.LineSpacing)
			if hover(pCursor, b.Size, b.position, img, 1) {
				text.Draw(img, line, b.font, tx, ty, b.colorTextHover)
			} else {
				text.Draw(img, line, b.font, tx, ty, b.colorText)
			}
		}
	} else {
		ebitenutil.DrawRect(img, 0, 0+float64(b.radius), b.Size.X, b.Size.Y-2*float64(b.radius), b.colorButton)
		ebitenutil.DrawRect(img, 0+float64(b.radius), 0, b.Size.X-2*float64(b.radius), b.Size.Y, b.colorButton)
		ebitenutil.DrawCircle(img, 0+float64(b.radius), 0+float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+b.Size.X-float64(b.radius), 0+float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+float64(b.radius), 0+b.Size.Y-float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+b.Size.X-float64(b.radius), 0+b.Size.Y-float64(b.radius), float64(b.radius), b.colorButton)
		// rounded corners: draw base shapes
		ebitenutil.DrawRect(img, 0, 0+float64(b.radius), b.Size.X, float64(height)-2*float64(b.radius), b.colorButton)
		ebitenutil.DrawRect(img, 0+float64(b.radius), 0, b.Size.X-2*float64(b.radius), float64(height), b.colorButton)
		ebitenutil.DrawCircle(img, 0+float64(b.radius), 0+float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+b.Size.X-float64(b.radius), 0+float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+float64(b.radius), 0+float64(height)-float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+b.Size.X-float64(b.radius), 0+float64(height)-float64(b.radius), float64(b.radius), b.colorButton)
		if hover(pCursor, b.Size, b.position, img, 1) {
			ebitenutil.DrawRect(img, b.borderSize, b.borderSize+float64(b.radius), b.Size.X-b.borderSize*2, float64(height)-b.borderSize*2-2*float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawRect(img, b.borderSize+float64(b.radius), b.borderSize, b.Size.X-b.borderSize*2-2*float64(b.radius), float64(height)-b.borderSize*2, b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize+float64(b.radius), b.borderSize+float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize+b.Size.X-b.borderSize*2-float64(b.radius), b.borderSize+float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize+float64(b.radius), b.borderSize+float64(height)-b.borderSize*2-float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize+b.Size.X-b.borderSize*2-float64(b.radius), b.borderSize+float64(height)-b.borderSize*2-float64(b.radius), float64(b.radius), b.colorButtonHover)
		}
		// draw text lines centered (account for LineSpacing)
		totalTextHeight := lineHeight*len(lines) + extraSpacing
		startY := (height-totalTextHeight)/2 + ascent
		for i, line := range lines {
			d := text.BoundString(b.font, line)
			lineW := d.Max.X - d.Min.X
			tx := int(float64(width)/2.0) - lineW/2
			ty := startY + i*(lineHeight+b.LineSpacing)
			if hover(pCursor, b.Size, b.position, img, 1) {
				text.Draw(img, line, b.font, tx, ty, b.colorTextHover)
			} else {
				text.Draw(img, line, b.font, tx, ty, b.colorText)
			}
		}
	}
	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Translate(b.position.X, b.position.Y)
	b.img = img
	screen.DrawImage(img, ot)
}

func (b *Button) DrawAt(screen *ebiten.Image, pos gfx.Point) {
	oldPos := b.position
	b.position = pos
	if b.Size.X <= 0 {
		b.Size.X = 100
	}
	if b.Size.Y <= 0 {
		b.Size.Y = 30
	}
	b.Draw(screen)
	b.position = oldPos
}

func (b *Button) Input() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pCursor := gfx.Point{X: float64(x), Y: float64(y)}
		if hover(pCursor, b.Size, b.position, b.img, 1) {
			b.Execute = true
		}
	}
}

func (b *Button) SetRadius(radius int) {
	if radius > int(b.Size.X/2) {
		b.radius = int(b.Size.X/2) - int(b.borderSize)
	}
	if radius > int(b.Size.Y/2) {
		b.radius = int(b.Size.Y/2) - int(b.borderSize)
	} else {
		b.radius = radius
	}
}

func (b *Button) SetColor(colorButton color.RGBA) {
	b.colorButton = colorButton
}

func (b *Button) SetColorHover(colorButtonHover color.RGBA) {
	b.colorButtonHover = colorButtonHover
}

func (b *Button) SetText(text string) {
	b.text = text
}

func (b *Button) SetFont(path string) {
	tt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fontParsed, erre := truetype.Parse(tt)
	if erre != nil {
		panic(erre)
	}
	b.fontParsed = fontParsed
	fontFace := truetype.NewFace(b.fontParsed, &truetype.Options{
		Size:    b.fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	b.font = fontFace
}

func (b *Button) SetFontSize(size float64) {
	b.fontSize = size
	if b.fontParsed == nil {
		fmt.Println("Warning - as long as you haven't changed the font, it can't change its size")
	} else {
		fontFace := truetype.NewFace(b.fontParsed, &truetype.Options{
			Size:    b.fontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		b.font = fontFace
	}
}

func (b *Button) SetPosition(position gfx.Point) {
	b.position = position
}

func (b *Button) SetColorText(colorText color.RGBA) {
	b.colorText = colorText
}

func (b *Button) SetColorTextHover(colorTextHover color.RGBA) {
	b.colorTextHover = colorTextHover
}

func (b *Button) SetBorderSize(size float64) {
	b.borderSize = size
}

type SpriteButton struct {
	position   gfx.Point
	scale      float64
	imgDefault *ebiten.Image
	imgHover   *ebiten.Image
	imgClicked *ebiten.Image
	Execute    bool
}

func NewSpriteButton(position gfx.Point, pathImgDefault, pathImgHover, pathImgClicked string) SpriteButton {
	imgDefault, _, err := ebitenutil.NewImageFromFile(pathImgDefault)
	imgHover, _, errs := ebitenutil.NewImageFromFile(pathImgHover)
	imgClicked, _, erres := ebitenutil.NewImageFromFile(pathImgClicked)
	if err != nil || errs != nil || erres != nil {
		log.Fatalf("Failed to load image for sprite button :\n%s\n%s\n%s", pathImgDefault, pathImgHover, pathImgClicked)
	}
	xDef, yDef := imgDefault.Size()
	xHov, yHov := imgHover.Size()
	xCli, yCli := imgClicked.Size()
	if xDef != xHov || yDef != yHov || xDef != xCli || yDef != yCli {
		log.Fatalf("ERROR with spriteButton the 3 image don't have same dimension")
	}
	return SpriteButton{
		position,
		1,
		imgDefault,
		imgHover,
		imgClicked,
		false,
	}
}

func (b *SpriteButton) Draw(screen *ebiten.Image) {
	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Scale(b.scale, b.scale)
	ot.GeoM.Translate(b.position.X, b.position.Y)
	x, y := ebiten.CursorPosition()
	pCursor := gfx.Point{X: float64(x), Y: float64(y)}
	xx, yy := b.imgDefault.Size()
	size := gfx.Point{X: float64(xx) * b.scale, Y: float64(yy) * b.scale}
	if hover(pCursor, size, b.position, b.imgDefault, b.scale) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			screen.DrawImage(b.imgClicked, ot)
		} else {
			screen.DrawImage(b.imgHover, ot)
		}
	} else {
		screen.DrawImage(b.imgDefault, ot)
	}
}

func (b *SpriteButton) Input() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pCursor := gfx.Point{X: float64(x), Y: float64(y)}
		xx, yy := b.imgDefault.Size()
		size := gfx.Point{X: float64(xx) * b.scale, Y: float64(yy) * b.scale}
		if hover(pCursor, size, b.position, b.imgDefault, b.scale) {
			b.Execute = true
		}
	}
}

func (b *SpriteButton) SetScale(scale float64) {
	if scale > 1 || scale < 0 {
		b.scale = 1.0
	} else {
		b.scale = scale
	}
}
