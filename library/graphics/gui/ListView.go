package gui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/vad3l/gold/library/graphics"
)

type ListView struct {
	Position      Point
	Size          Point
	Items         []Widget
	ScrollOffset  float64
	ItemSpacing   float64
	Background    color.Color
	BorderColor   color.Color
	SelectedIndex int
	ItemHeight    float64
}

func NewListView(position, size Point) *ListView {
	return &ListView{
		Position:      position,
		Size:          size,
		ScrollOffset:  0,
		ItemSpacing:   10,
		ItemHeight:    40,
		Background:    color.RGBA{220, 220, 220, 255},
		BorderColor:   color.RGBA{0, 0, 0, 255},
		SelectedIndex: -1,
	}
}

func (lv *ListView) Add(item Widget) {
	lv.Items = append(lv.Items, item)
}

func (lv *ListView) Input() {
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		lv.ScrollOffset -= wheelY * 20
	}

	if lv.ScrollOffset < 0 {
		lv.ScrollOffset = 0
	}

	maxScroll := float64(len(lv.Items))*(lv.ItemHeight+lv.ItemSpacing) - lv.Size.Y
	if maxScroll < 0 {
		maxScroll = 0
	}
	if lv.ScrollOffset > maxScroll {
		lv.ScrollOffset = maxScroll
	}

	x, y := ebiten.CursorPosition()
	pCursor := Point{float64(x), float64(y)}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for i := range lv.Items {
			itemY := lv.Position.Y + float64(i)*(lv.ItemHeight+lv.ItemSpacing) - lv.ScrollOffset
			if pCursor.Y >= itemY && pCursor.Y <= itemY+lv.ItemHeight &&
				pCursor.X >= lv.Position.X && pCursor.X <= lv.Position.X+lv.Size.X {
				lv.SelectedIndex = i
				break
			}
		}
	}
}

func (lv *ListView) Draw(screen *ebiten.Image) {
	width := int(math.Max(lv.Size.X, 1))
	height := int(math.Max(lv.Size.Y, 1))

	bg := ebiten.NewImage(width, height)
	bg.Fill(lv.Background)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(lv.Position.X, lv.Position.Y)
	screen.DrawImage(bg, op)

	startY := lv.Position.Y - lv.ScrollOffset
	for i, item := range lv.Items {
		itemY := startY + float64(i)*(lv.ItemHeight+lv.ItemSpacing)
		if itemY+lv.ItemHeight < lv.Position.Y {
			continue
		}
		if itemY > lv.Position.Y+lv.Size.Y {
			break
		}

		if i == lv.SelectedIndex {
			sel := ebiten.NewImage(width, int(lv.ItemHeight))
			sel.Fill(color.RGBA{180, 200, 255, 255})
			opSel := &ebiten.DrawImageOptions{}
			opSel.GeoM.Translate(lv.Position.X, itemY)
			screen.DrawImage(sel, opSel)
		}

		if drawable, ok := item.(interface{ DrawAt(*ebiten.Image, Point) }); ok {
			drawable.DrawAt(screen, Point{lv.Position.X + 10, itemY})
		}
	}

	border := ebiten.NewImage(width, height)
	for x := 0; x < width; x++ {
		border.Set(x, 0, lv.BorderColor)
		border.Set(x, height-1, lv.BorderColor)
	}
	for y := 0; y < height; y++ {
		border.Set(0, y, lv.BorderColor)
		border.Set(width-1, y, lv.BorderColor)
	}
	opBorder := &ebiten.DrawImageOptions{}
	opBorder.GeoM.Translate(lv.Position.X, lv.Position.Y)
	screen.DrawImage(border, opBorder)
}
