package appgui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"image/color"
)

func Encodebox() fyne.CanvasObject {
	encodetext := widget.NewMultiLineEntry()
	encodetext.Wrapping = fyne.TextWrapWord
	encodetext.SetPlaceHolder("Encode")
	textContainer := widget.NewScrollContainer(encodetext)
	textContainer.SetMinSize(fyne.NewSize(1040, 140))
	box := widget.NewVBox(
		widget.NewLabel("Encode"),
		canvas.NewLine(color.Black),
		textContainer,
		canvas.NewLine(color.Black),
	)
	return box
}

func MakeSplitTab() fyne.CanvasObject {

	decodetext := widget.NewMultiLineEntry()
	decodetext.BaseWidget.Resize(fyne.NewSize(12, 12))
	decodetext.Wrapping = fyne.TextWrapWord
	decodetext.SetPlaceHolder("Decode")

	right := widget.NewVBox(
		Encodebox(),
		widget.NewButton("Button", func() { fmt.Println("button tapped!") }),
		decodetext,
	)
	//result := widget.NewHSplitContainer(widget.NewVScrollContainer(left), right)
	return right
}
