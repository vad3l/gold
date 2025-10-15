package gui

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	gfx "github.com/vad3l/gold/library/graphics"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type ListView struct {
	Position      gfx.Point
	Size          gfx.Point
	Items         []Widget
	ScrollOffset  float64
	ItemSpacing   float64
	Background    color.Color
	BorderColor   color.Color
	SelectedIndex int
	ItemHeight    float64
	// scrollbar
	ShowScrollBar  bool
	ScrollBarWidth float64
	thumbDragging  bool
}

func NewListView(position, size gfx.Point) *ListView {
	return &ListView{
		Position:       position,
		Size:           size,
		ScrollOffset:   0,
		ItemSpacing:    10,
		ItemHeight:     40,
		Background:     color.RGBA{220, 220, 220, 255},
		BorderColor:    color.RGBA{0, 0, 0, 255},
		SelectedIndex:  -1,
		ShowScrollBar:  true,
		ScrollBarWidth: 12,
	}
}

func (lv *ListView) Add(item Widget) {
	lv.Items = append(lv.Items, item)
}

func (lv *ListView) Input() {
	// Wheel scrolling
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		// multiply to get a comfortable speed
		lv.ScrollOffset -= wheelY * 20
	}

	// Clamp scroll to valid range
	// Compute content height and clamp scroll to valid range
	contentHeight := float64(len(lv.Items))*(lv.ItemHeight+lv.ItemSpacing) - lv.ItemSpacing
	if contentHeight < 0 {
		contentHeight = 0
	}
	maxScroll := contentHeight - lv.Size.Y
	if maxScroll < 0 {
		maxScroll = 0
	}
	if lv.ScrollOffset < 0 {
		lv.ScrollOffset = 0
	}
	if lv.ScrollOffset > maxScroll {
		lv.ScrollOffset = maxScroll
	}

	// Click handling: scrollbar interaction first
	// Scrollbar is drawn inside the right edge of the list
	trackX := lv.Position.X + lv.Size.X - lv.ScrollBarWidth
	trackY := lv.Position.Y
	trackW := lv.ScrollBarWidth
	trackH := lv.Size.Y
	// compute thumb size and position
	var thumbH float64
	if contentHeight <= 0 || contentHeight <= lv.Size.Y {
		thumbH = lv.Size.Y
	} else {
		ratio := lv.Size.Y / contentHeight
		thumbH = ratio * lv.Size.Y
		if thumbH < 20 {
			thumbH = 20
		}
	}
	var thumbY float64
	if maxScroll <= 0 {
		thumbY = trackY
	} else {
		thumbY = trackY + (lv.ScrollOffset/maxScroll)*(trackH-thumbH)
	}

	// Handle scrollbar mouse interactions
	x, y := ebiten.CursorPosition()
	pCursor := gfx.Point{X: float64(x), Y: float64(y)}
	// mouse press: start dragging if on thumb
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if pCursor.X >= trackX && pCursor.X <= trackX+trackW && pCursor.Y >= thumbY && pCursor.Y <= thumbY+thumbH {
			lv.thumbDragging = true
			return
		}
		// click on track (page up/down)
		if pCursor.X >= trackX && pCursor.X <= trackX+trackW && pCursor.Y >= trackY && pCursor.Y <= trackY+trackH {
			// click above thumb => page up
			if pCursor.Y < thumbY {
				lv.ScrollOffset -= lv.Size.Y
			} else if pCursor.Y > thumbY+thumbH {
				lv.ScrollOffset += lv.Size.Y
			}
			if lv.ScrollOffset < 0 {
				lv.ScrollOffset = 0
			}
			if lv.ScrollOffset > maxScroll {
				lv.ScrollOffset = maxScroll
			}
			return
		}
	}

	// dragging
	if lv.thumbDragging && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// compute new thumb center relative position
		rel := pCursor.Y - trackY - thumbH/2
		if rel < 0 {
			rel = 0
		}
		if rel > trackH-thumbH {
			rel = trackH - thumbH
		}
		if trackH-thumbH > 0 {
			lv.ScrollOffset = (rel / (trackH - thumbH)) * maxScroll
		} else {
			lv.ScrollOffset = 0
		}
		return
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// stop dragging when released
		if lv.thumbDragging {
			lv.thumbDragging = false
			return
		}

		// Make sure click is inside list bounds (excluding scrollbar area)
		if pCursor.X < lv.Position.X || pCursor.X > lv.Position.X+lv.Size.X-lv.ScrollBarWidth || pCursor.Y < lv.Position.Y || pCursor.Y > lv.Position.Y+lv.Size.Y {
			return
		}

		for i := range lv.Items {
			itemY := lv.Position.Y + float64(i)*(lv.ItemHeight+lv.ItemSpacing) - lv.ScrollOffset
			// Only consider visible items
			if itemY+lv.ItemHeight < lv.Position.Y {
				continue
			}
			if itemY > lv.Position.Y+lv.Size.Y {
				break
			}
			if pCursor.Y >= itemY && pCursor.Y <= itemY+lv.ItemHeight {
				lv.SelectedIndex = i
				// If the item is a Button (pointer), trigger its Execute and print its text
				switch it := lv.Items[i].(type) {
				case *Button:
					it.Execute = true
					fmt.Println(it.text)
				case *SpriteButton:
					it.Execute = true
				default:
					// For other widgets, call their Input method so they can react if they use absolute positions
					lv.Items[i].Input()
				}
				break
			}
		}
	}
}

func (lv *ListView) Draw(screen *ebiten.Image) {
	width := int(math.Max(lv.Size.X, 1))
	height := int(math.Max(lv.Size.Y, 1))

	// Render into a viewport image equal to the list size so content is clipped
	viewport := ebiten.NewImage(width, height)
	viewport.Fill(lv.Background)

	startY := -lv.ScrollOffset
	for i, item := range lv.Items {
		itemY := startY + float64(i)*(lv.ItemHeight+lv.ItemSpacing)
		// Skip items that are completely above the viewport
		if itemY+lv.ItemHeight < 0 {
			continue
		}
		// Stop when item starts below the viewport
		if itemY > lv.Size.Y {
			break
		}

		// Draw item background and content with padding to resemble a navigator list
		padding := 10.0
		itemX := padding
		// leave space for scrollbar on the right
		itemW := lv.Size.X - lv.ScrollBarWidth - padding*2

		// base background
		bg := color.RGBA{245, 245, 245, 255}
		hoverBg := color.RGBA{220, 235, 255, 255}
		itemBg := bg
		// check hover relative to viewport: mouse pos is global, convert to local
		mx, my := ebiten.CursorPosition()
		mxF, myF := float64(mx)-lv.Position.X, float64(my)-lv.Position.Y
		if mxF >= itemX && mxF <= itemX+itemW && myF >= itemY && myF <= itemY+lv.ItemHeight {
			itemBg = hoverBg
		}
		// draw item background
		ebitenutil.DrawRect(viewport, itemX, itemY, itemW, lv.ItemHeight, itemBg)

		// small icon on left
		iconX := itemX + 8
		iconY := itemY + lv.ItemHeight/2
		ebitenutil.DrawCircle(viewport, iconX, iconY, lv.ItemHeight*0.25, color.RGBA{100, 140, 220, 255})

		// text
		var txt string
		var fnt font.Face = basicfont.Face7x13
		switch it := item.(type) {
		case *Button:
			txt = it.text
			if it.font != nil {
				fnt = it.font
			}
		default:
			txt = "item"
		}
		textX := int(itemX + 8 + lv.ItemHeight*0.5 + 8)
		textY := int(itemY + lv.ItemHeight/2 + 6)
		text.Draw(viewport, txt, fnt, textX, textY, color.RGBA{20, 20, 20, 255})

		// divider line
		ebitenutil.DrawRect(viewport, itemX, itemY+lv.ItemHeight-1, itemW, 1, color.RGBA{220, 220, 220, 255})
	}

	// Blit the viewport to the screen at the ListView position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(lv.Position.X, lv.Position.Y)
	screen.DrawImage(viewport, op)

	// Draw scrollbar (track + thumb) on the right side of the viewport
	if lv.ShowScrollBar {
		trackW := int(math.Max(lv.ScrollBarWidth, 1))
		trackX := int(lv.Position.X + lv.Size.X - lv.ScrollBarWidth)
		trackY := int(lv.Position.Y)
		trackH := int(lv.Size.Y)

		// compute content height and thumb size/pos (same as in Input)
		contentHeight := float64(len(lv.Items))*(lv.ItemHeight+lv.ItemSpacing) - lv.ItemSpacing
		if contentHeight < 0 {
			contentHeight = 0
		}
		var thumbH float64
		if contentHeight <= 0 || contentHeight <= lv.Size.Y {
			thumbH = lv.Size.Y
		} else {
			ratio := lv.Size.Y / contentHeight
			thumbH = ratio * lv.Size.Y
			if thumbH < 20 {
				thumbH = 20
			}
		}
		maxScroll := contentHeight - lv.Size.Y
		if maxScroll < 0 {
			maxScroll = 0
		}
		var thumbY float64
		if maxScroll <= 0 {
			thumbY = lv.Position.Y
		} else {
			thumbY = lv.Position.Y + (lv.ScrollOffset/maxScroll)*(lv.Size.Y-thumbH)
		}

		// draw track
		trackImg := ebiten.NewImage(trackW, trackH)
		trackImg.Fill(color.RGBA{200, 200, 200, 255})
		opTrack := &ebiten.DrawImageOptions{}
		opTrack.GeoM.Translate(float64(trackX), float64(trackY))
		screen.DrawImage(trackImg, opTrack)

		// draw thumb
		thumbImg := ebiten.NewImage(trackW-4, int(math.Max(thumbH, 4)))
		thumbImg.Fill(color.RGBA{150, 150, 150, 255})
		opThumb := &ebiten.DrawImageOptions{}
		opThumb.GeoM.Translate(float64(trackX+2), thumbY)
		screen.DrawImage(thumbImg, opThumb)
	}

	// Draw border
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
