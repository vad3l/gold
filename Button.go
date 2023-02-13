package main

import (
	"image/color"
	_ "image/png"
	"io/ioutil"
	"fmt"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct{
	size	Point
	position	Point
	colorButton	color.RGBA
	colorButtonHover	color.RGBA
	borderSize	float64

	text	string
	font	font.Face
	fontSize	float64
	fontParsed	*truetype.Font
	colorText	color.RGBA
	colorTextHover	color.RGBA

	radius	int
	execute	func(g *SceneManager)
	img	*ebiten.Image
}

func NewButton (size, position Point, text string, execute func(g *SceneManager)) Button{

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
		execute,
		nil,
	}
}

func (b *Button) Draw (screen *ebiten.Image) {
	img := ebiten.NewImage(int(b.size.x), int(b.size.y))
	b.img = img
	x, y := ebiten.CursorPosition()
	pCursor := Point{float64(x),float64(y)}
	
	fontDimension := text.BoundString(b.font,b.text)
	height := fontDimension.Max.Y * 2
	if (height == 0){
		height = int(b.fontSize / 2 ) - 3 
	}
	tx := int(b.size.x / 2) - (fontDimension.Max.X /2 )
	ty := int(b.size.y / 2) + height

	if (b.radius == 0){
		ebitenutil.DrawRect(img, 0, 0, b.size.x, b.size.y, b.colorButton)
		if b.hover(pCursor) {
			ebitenutil.DrawRect(img, b.borderSize, b.borderSize, b.size.x-b.borderSize*2, b.size.y-b.borderSize*2, b.colorButtonHover)
			text.Draw(img, b.text, b.font, tx, ty, b.colorTextHover)
		} else {
			text.Draw(img, b.text, b.font, tx, ty, b.colorText)
		}
	}else{
		ebitenutil.DrawRect(img, 0, 0 + float64(b.radius), b.size.x, b.size.y - 2*float64(b.radius), b.colorButton)
		ebitenutil.DrawRect(img, 0 + float64(b.radius), 0, b.size.x - 2*float64(b.radius), b.size.y, b.colorButton)

		ebitenutil.DrawCircle(img, 0 + float64(b.radius), 0 + float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0 + b.size.x - float64(b.radius), 0 + float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0 + float64(b.radius), 0 + b.size.y - float64(b.radius), float64(b.radius), b.colorButton)
		ebitenutil.DrawCircle(img, 0 + b.size.x - float64(b.radius), 0 + b.size.y - float64(b.radius), float64(b.radius), b.colorButton)
		if b.hover(pCursor) {
			ebitenutil.DrawRect(img, b.borderSize, b.borderSize  + float64(b.radius), b.size.x-b.borderSize*2, b.size.y-b.borderSize*2 - 2*float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawRect(img, b.borderSize + float64(b.radius), b.borderSize, b.size.x-b.borderSize*2 - 2*float64(b.radius), b.size.y-b.borderSize*2, b.colorButtonHover)

			ebitenutil.DrawCircle(img, b.borderSize + float64(b.radius), b.borderSize + float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize + b.size.x-b.borderSize*2 - float64(b.radius), b.borderSize + float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize + float64(b.radius), b.borderSize + b.size.y-b.borderSize*2 - float64(b.radius), float64(b.radius), b.colorButtonHover)
			ebitenutil.DrawCircle(img, b.borderSize + b.size.x-b.borderSize*2 - float64(b.radius), b.borderSize + b.size.y-b.borderSize*2 - float64(b.radius), float64(b.radius), b.colorButtonHover)
			text.Draw(img, b.text, b.font, tx, ty, b.colorTextHover)
		}else{
			text.Draw(img, b.text, b.font, tx, ty, b.colorText)
		}
	}

	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Translate(b.position.x, b.position.y)
	b.img = img
	screen.DrawImage(img, ot)
}

func (b *Button) Input (g *SceneManager) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pCursor := Point{float64(x),float64(y)}
		if b.hover(pCursor) {
			b.execute(g)
		}
	}
}

func (b *Button) hover (pCursor Point) bool  {
	/*
	dx := pCursor.x - circle.Position.x
	dy := pCursor.y - circle.Position.y
	distance := math.Sqrt(dx*dx + dy*dy)

	if distance <= float64(circle.Radius) {
		return true
	}*/
	if ((pCursor.x <= (b.size.x + b.position.x) && pCursor.x >= b.position.x) && (pCursor.y <= (b.size.y + b.position.y) && pCursor.y >= b.position.y)){
		c := b.img.At(int(pCursor.x) - int(b.position.x), int(pCursor.y)- int(b.position.y) ).(color.RGBA)	
		if c.A > 0 {
			return true
		}
		return false
	}
	return false
}

func (b *Button) setRadius (radius int) {
	if ( radius > int(b.size.x / 2 ) ){
		b.radius = int(b.size.x / 2 ) - int(b.borderSize)
	}
	if ( radius > int(b.size.y / 2 ) ) {
		b.radius = int(b.size.y / 2 ) - int(b.borderSize)
	}else{
		b.radius = radius
	}
}

func (b *Button) setColor (colorButton color.RGBA) {
	b.colorButton = colorButton
}

func (b *Button) setColorHover (colorButtonHover color.RGBA) {
	b.colorButtonHover = colorButtonHover
}

func (b *Button) setText (text string) {
	b.text = text
}

func (b *Button) setFont (fontName string) {
	tt, err := ioutil.ReadFile("./data/font/"+fontName)
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

func (b *Button) setFontSize (size float64) {
	b.fontSize = size
	if (b.fontParsed == nil){
		fmt.Println("Warning - as long as you haven't changed the font, it can't change its size")
	}else{
		fontFace := truetype.NewFace(b.fontParsed, &truetype.Options{
			Size:    b.fontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		b.font = fontFace
	}
	
	
}

func (b *Button) setColorText (colorText color.RGBA) {
	b.colorText = colorText
}

func (b *Button) setColorTextHover (colorTextHover color.RGBA) {
	b.colorTextHover = colorTextHover
}

func (b *Button) setBorderSize (size float64) {
	b.borderSize = size
}

/*******************
*
*	Sprite Button
*
*///////////////////


type SpriteButton struct{
	position	Point
	scale	float64

	imgDefault	*ebiten.Image
	imgHover	*ebiten.Image
	imgClicked	*ebiten.Image
	
	execute	func(g *SceneManager)
}

func NewSpriteButton (position Point, pathImgDefault, pathImgHover, pathImgClicked string, execute func(g *SceneManager)) SpriteButton{
	imgDefault, _, err := ebitenutil.NewImageFromFile("data/image/button/"+pathImgDefault)
	imgHover, _, errs := ebitenutil.NewImageFromFile("data/image/button/"+pathImgHover)
	imgClicked, _, erres := ebitenutil.NewImageFromFile("data/image/button/"+pathImgClicked)
	if (err != nil  || errs != nil || erres != nil ){
		log.Fatalf("Failed to load image for sprite button :\n%s\n%s\n%s",pathImgDefault,pathImgHover,pathImgClicked)
	}

	xDef,yDef := imgDefault.Size()
	xHov,yHov := imgHover.Size()
	xCli,yCli := imgClicked.Size()
	if (xDef != xHov || yDef != yHov || xDef != xCli || yDef != yCli ){
		log.Fatalf("ERROR with spriteButton the 3 image don't have same dimension")
	}

	return SpriteButton{
		position,
		1,
		imgDefault,
		imgHover,
		imgClicked,
		execute,
	}
}

func (b *SpriteButton) Draw (screen *ebiten.Image) {
	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Scale(b.scale,b.scale)
	ot.GeoM.Translate(b.position.x, b.position.y)
	

	x, y := ebiten.CursorPosition()
	pCursor := Point{float64(x),float64(y)}
	if b.hover(pCursor) {
		if (ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)){
			screen.DrawImage(b.imgClicked, ot)
		}else{
			screen.DrawImage(b.imgHover, ot)
		}
	}else {
		screen.DrawImage(b.imgDefault, ot)
	}
}
func (b *SpriteButton) Input (g *SceneManager) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pCursor := Point{float64(x),float64(y)}
		if b.hover(pCursor) {
			b.execute(g)
		}
	}
}

func (b *SpriteButton) hover (pCursor Point) bool  {
	x,y := b.imgDefault.Size()
	size := Point{float64(x)*b.scale,float64(y)*b.scale}
	if ((pCursor.x <= (size.x + b.position.x) && pCursor.x >= b.position.x) && (pCursor.y <= (size.y + b.position.y) && pCursor.y >= b.position.y)){	
		c := b.imgHover.At(int(pCursor.x/b.scale) - int(b.position.x /b.scale), int(pCursor.y/b.scale)- int(b.position.y /b.scale) ).(color.RGBA)	
		if c.A > 0 {
			return true
		}
		return false
	}

	
	return false
}

func (b *SpriteButton) setScale (scale float64) {
	if (scale > 1 || scale < 0){
		b.scale = 1.0
	}else{
		b.scale = scale
	}
}