package appgui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func makeButtonList(count int) []fyne.CanvasObject {
	var items []fyne.CanvasObject
	for i := 1; i <= count; i++ {
		index := i // capture
		items = append(items, widget.NewButton(fmt.Sprintf("Button %d", index), func() {
			fmt.Println("Tapped", index)
		}))
	}

	return items
}

func Button() fyne.CanvasObject {
	bt := widget.NewButton("URL编码UTF-8", func() {
		fmt.Println("test")
	})
	return bt
}

func test() {
	fmt.Println("123")

}
func Bottonfrom() fyne.CanvasObject {
	hlist := makeButtonList(20)
	//vlist := makeButtonList(50)

	horiz := widget.NewHScrollContainer(widget.NewHBox(widget.NewLabel("加密："), hlist[0], hlist[1], hlist[2], hlist[3], hlist[4], hlist[5]))
	//vert := widget.NewVScrollContainer(widget.NewVBox(vlist...))
	box := widget.NewGroupWithScroller("123", horiz)
	return box
}
