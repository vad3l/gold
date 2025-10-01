package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	. "github.com/vad3l/gold/library/graphics"
)

type Slider struct {
	size       Point
	position   Point
	barColor   color.RGBA
	thumbColor color.RGBA
	min, max   float64
	value      float64
	dragging   bool
}

func NewSlider(size, position Point, min, max, value float64) Slider {
	return Slider{
		size:       size,
		position:   position,
		barColor:   color.RGBA{0xff, 0x00, 0x00, 0xff},
		thumbColor: color.RGBA{0x00, 0xff, 0x00, 0xff},
		min:        min,
		max:        max,
		value:      value,
		dragging:   false,
	}
}

func (s *Slider) Draw(screen *ebiten.Image) {
	// Draw bar
	barY := s.position.Y + s.size.Y/2 - 3
	ebitenutil.DrawRect(screen, s.position.X, barY, s.size.X, 6, s.barColor)

	// Draw thumb
	thumbX := s.position.X + ((s.value-s.min)/(s.max-s.min))*s.size.X
	thumbY := s.position.Y + s.size.Y/2
	ebitenutil.DrawCircle(screen, thumbX, thumbY, s.size.Y/2, s.thumbColor)
}

func (s *Slider) Input() {
	x, y := ebiten.CursorPosition()
	cursor := Point{float64(x), float64(y)}
	thumbX := s.position.X + ((s.value-s.min)/(s.max-s.min))*s.size.X
	thumbY := s.position.Y + s.size.Y/2

	// Start dragging if mouse pressed on thumb
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		dx := cursor.X - thumbX
		dy := cursor.Y - thumbY
		r := s.size.Y / 2
		if dx*dx+dy*dy <= r*r {
			s.dragging = true
		}
	}
	// Stop dragging
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		s.dragging = false
	}
	// Update value while dragging
	if s.dragging {
		rel := cursor.X - s.position.X
		if rel < 0 {
			rel = 0
		}
		if rel > s.size.X {
			rel = s.size.X
		}
		s.value = s.min + (rel/s.size.X)*(s.max-s.min)
	}
}

func (s *Slider) SetBarColor(color color.RGBA) {
	s.barColor = color
}

func (s *Slider) SetThumbColor(color color.RGBA) {
	s.thumbColor = color
}

// Optionally, add getters if you want to access value from outside the package
func (s *Slider) Value() float64 {
	return s.value
}
