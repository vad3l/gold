package gui

import (
	"fmt"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"

	"github.com/golang/freetype/truetype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	. "github.com/vad3l/gold/library/graphics"
)

type Button struct {
	Size             Point
	position         Point
	colorButton      color.RGBA
	colorButtonHover color.RGBA
	borderSize       float64

	text           string
	font           font.Face
	fontSize       float64
	fontParsed     *truetype.Font
	colorText      color.RGBA
	colorTextHover color.RGBA

	radius  int
	Execute bool
	img     *ebiten.Image
}

func NewButton(size, position Point, text string) Button {

	return Button{
		size,
		position,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		10.0,
		text,
		basicfont.Face7x13,
		12,
		nil,
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		0,
		false,
		nil,
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(b.Size.X), int(b.Size.Y))
	b.img = img
	x, y := ebiten.CursorPosition()
	pCursor := Point{float64(x), float64(y)}

	fontDimension := text.BoundString(b.font, b.text)
	height := fontDimension.Max.Y * 2
	if height == 0 {
		height = int(b.fontSize/2) - 3
	}
	tx := int(b.Size.X/2) - (fontDimension.Max.X / 2)
	ty := int(b.Size.Y/2) + height

	if b.radius == 0 {
		ebitenutil.DrawRect(img, 0, 0, b.Size.X, b.Size.Y, b.colorButton)
		if hover(pCursor, b.Size, b.position, img, 1) {
			ebitenutil.DrawRect(img, b.borderSize, b.borderSize, b.Size.X-b.borderSize*2, b.Size.Y-b.borderSize*2, b.colorButtonHover)
			text.Draw(img, b.text, b.font, tx, ty, b.colorTextHover)
		} else {
			text.Draw(img, b.text, b.font, tx, ty, b.colorText)
		}
	} else {
		ebitenutil.DrawRect(img, 0, 0+float64(b.radius), b.Size.X, b.Size.Y-2*float64(b.radius), b.colorButton)
		ebitenutil.DrawRect(img, 0+float64(b.radius), 0, b.Size.X-2*float64(b.radius), b.Size.Y, b.colorButton)

		ebitenutil.DrawCircle(img, 0+float64(b.radius), 0+float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+b.Size.X-float64(b.radius), 0+float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+float64(b.radius), 0+b.Size.Y-float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0+b.Size.X-float64(b.radius), 0+b.Size.Y-float64(b.radius), float64(b.radius), b.colorButton)
		if hover(pCursor, b.Size, b.position, img, 1) {
			ebitenutil.DrawRect(img, b.borderSize, b.borderSize+float64(b.radius), b.Size.X-b.borderSize*2, b.Size.Y-b.borderSize*2-2*float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawRect(img, b.borderSize+float64(b.radius), b.borderSize, b.Size.X-b.borderSize*2-2*float64(b.radius), b.Size.Y-b.borderSize*2, b.colorButtonHover)

			ebitenutil.DrawCircle(img, b.borderSize+float64(b.radius), b.borderSize+float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize+b.Size.X-b.borderSize*2-float64(b.radius), b.borderSize+float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize+float64(b.radius), b.borderSize+b.Size.Y-b.borderSize*2-float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize+b.Size.X-b.borderSize*2-float64(b.radius), b.borderSize+b.Size.Y-b.borderSize*2-float64(b.radius), float64(b.radius), b.colorButtonHover)
			text.Draw(img, b.text, b.font, tx, ty, b.colorTextHover)
		} else {
			text.Draw(img, b.text, b.font, tx, ty, b.colorText)
		}
	}

	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Translate(b.position.X, b.position.Y)
	b.img = img
	screen.DrawImage(img, ot)
}

func (b *Button) Input() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pCursor := Point{float64(x), float64(y)}
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

	// Create a font face with a specific size
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

func (b *Button) SetPosition(position Point) {
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

/*******************
*
*	Sprite Button
*
 */ //////////////////

type SpriteButton struct {
	position Point
	scale    float64

	imgDefault *ebiten.Image
	imgHover   *ebiten.Image
	imgClicked *ebiten.Image

	Execute bool
}

func NewSpriteButton(position Point, pathImgDefault, pathImgHover, pathImgClicked string) SpriteButton {
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
	pCursor := Point{float64(x), float64(y)}

	xx, yy := b.imgDefault.Size()
	size := Point{float64(xx) * b.scale, float64(yy) * b.scale}
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
		pCursor := Point{float64(x), float64(y)}
		xx, yy := b.imgDefault.Size()
		size := Point{float64(xx) * b.scale, float64(yy) * b.scale}
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
