package appgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
)

func BorderBody(width, hight int, body fyne.CanvasObject) *fyne.Container {
	top := HorizontalLine(width, 1)
	bottom := HorizontalLine(width, 1)
	left := VerticalLine(hight, 1)
	right := VerticalLine(hight, 1)
	borderset := layout.NewBorderLayout(top, bottom, left, right)

	return fyne.NewContainerWithLayout(borderset, top, bottom, left, right, body)
}

func HorizontalLine(length, strong int) fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.RGBA{0, 0, 0, 0xff})
	rect.SetMinSize(fyne.NewSize(length, strong))
	rect.Visible()
	return rect
}

func VerticalLine(length, strong int) fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.RGBA{0, 0, 0, 0xff})
	rect.SetMinSize(fyne.NewSize(strong, length))
	rect.Visible()
	return rect
}

func Textbox(text string) fyne.CanvasObject {
	inputbox := widget.NewMultiLineEntry()
	//inputbox.Wrapping = fyne.TextWrapWord
	inputbox.SetPlaceHolder(text)
	box := BorderBody(1040, 140, inputbox)
	return box
}
