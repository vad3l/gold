package gui

import (
	"image/color"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Slider struct {
	size	Point

	position	Point
	colorSlider	color.RGBA
	colorButton	color.RGBA
	
	circleCorner	bool

	cursorPosition	float64
	max	int
	img	*ebiten.Image
}

func NewSlider (size, position Point) Slider {
	if (size.X < size.Y) {
		size = Point{100,20}
	}
	maxInit := 68
	cursorPosition := maxInit/2
	return Slider{
		size,
		position,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		true,
		float64(cursorPosition),
		maxInit,
		nil,
	}
}

func (s *Slider) Draw (screen *ebiten.Image) {
	radius := s.size.Y/2
	img := ebiten.NewImage(int(s.size.X+ (4*radius)),int(s.size.Y+ (2*radius)))
	s.img = img

	x_new := float64(s.cursorPosition)
	if (s.max > 10) {
		x_new = (s.cursorPosition - 0) * (10 - 0) / (float64(s.max) - 0)
	}
	
	when0 := radius
	

	ebitenutil.DrawRect(img, radius, radius, s.size.X, s.size.Y, s.colorSlider)
	if (s.circleCorner){
		if (x_new == 0) {
			when0 += radius
		}
		ebitenutil.DrawCircle(img, radius, radius*2, radius, s.colorSlider)
		ebitenutil.DrawCircle(img, s.size.X+radius,  radius*2, radius, s.colorSlider)
		ebitenutil.DrawCircle(img,when0+(radius*2)*float64(x_new), s.size.Y , radius*2, s.colorButton)
	}else {
		if (x_new == 0) {
			when0 = 0
		}
		ebitenutil.DrawRect(img, (radius*2)*float64(x_new)-when0 , s.size.Y-radius*2, radius*4, radius*4, s.colorButton)
	}

	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Translate(s.position.X, s.position.Y)
	s.img = img
	screen.DrawImage(img, ot)
}

func (s *Slider) Input () {
	
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Println("click")
		x, y := ebiten.CursorPosition()
		pCursor := Point{float64(x),float64(y)}
		x_new := float64(s.cursorPosition)
		if (s.max > 10) {
			x_new = (s.cursorPosition - 0) * (10 - 0) / (float64(s.max) - 0)
		}
		sliderCursor := s.position.X-(s.size.Y/2)+(s.size.Y)*float64(x_new)
		if  hover(pCursor,Point{(s.size.Y*2),s.size.X*2},Point{sliderCursor,s.position.Y},s.img,1) {
			if (pCursor.X > sliderCursor + (s.size.Y) && x_new <10){
				s.cursorPosition = (x_new+1)*(float64(s.max)/10)
			} else if (x_new >= 0){
				s.cursorPosition = (x_new-1)*(float64(s.max)/10)
			}
		}
	}
}


func (s *Slider) SetColor (color color.RGBA) {
	s.colorSlider = color
}

func (s *Slider) SetColorButton (color color.RGBA) {
	s.colorButton= color
}