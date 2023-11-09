package ui

import (
	"fyne.io/fyne/v2"
	"ruhan.tech/pixl/apptype"
	"ruhan.tech/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
