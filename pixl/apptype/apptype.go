package apptype

import (
	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasConfig struct {
	Drawing        fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}
