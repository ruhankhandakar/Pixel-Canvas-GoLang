package ui

import (
	"fyne.io/fyne/v2"
	"ruhan.tech/pixl/apptype"
	"ruhan.tech/pixl/pxcanvas"
	"ruhan.tech/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
