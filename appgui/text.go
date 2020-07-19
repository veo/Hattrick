package appgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
)

type Textbox struct {
	InputText  *widget.Entry
	EncodeText *widget.Entry
	DecodeText *widget.Entry
}

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

func (t *Textbox) InputBox() fyne.CanvasObject {
	t.InputText = widget.NewMultiLineEntry()
	t.InputText.Wrapping = fyne.TextWrapWord
	t.InputText.SetPlaceHolder("inputText")
	box := BorderBody(1040, 140, t.InputText)
	return box
}
func (t *Textbox) EncodeBox() fyne.CanvasObject {
	t.EncodeText = widget.NewMultiLineEntry()
	t.EncodeText.Wrapping = fyne.TextWrapWord
	t.EncodeText.SetPlaceHolder("encodeText")
	box := BorderBody(1040, 140, t.EncodeText)
	return box
}

func (t *Textbox) DecodeBox() fyne.CanvasObject {
	t.DecodeText = widget.NewMultiLineEntry()
	t.DecodeText.Wrapping = fyne.TextWrapWord
	t.DecodeText.SetPlaceHolder("decodeText")
	box := BorderBody(1040, 140, t.DecodeText)
	return box
}
