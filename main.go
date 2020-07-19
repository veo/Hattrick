package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/veo/Hattrick/appgui"
)

func main() {
	app := app.New()
	app.SetIcon(appgui.Icon)
	app.Settings().SetTheme(&appgui.HatTrickTheme{})
	//app.Settings().SetTheme(theme.LightTheme())
	w := app.NewWindow("Hattrick")
	//w.SetContent(appgui.Bottonfrom())
	w.SetContent(widget.NewVBox(
		appgui.Textbox("Encode"),
		widget.NewButton("Button", func() { fmt.Println("button tapped!") }),
		appgui.Textbox("Result"),
		widget.NewButton("Button", func() { fmt.Println("button tapped!") }),
		appgui.Textbox("Decode"),
		widget.NewButton("Quit", func() { app.Quit() }),
	))
	w.Resize(fyne.NewSize(1100, 750))
	w.ShowAndRun()
}
