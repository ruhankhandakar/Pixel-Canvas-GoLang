package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"ruhan.tech/pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	ClickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}

	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  s,
	}
}

func NewSwatch(
	state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch),
) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		ClickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}

	swatch.ExtendBaseWidget(swatch)
	return swatch
}
