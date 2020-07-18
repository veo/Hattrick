package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"github.com/veo/Hattrick/appgui"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())
	w := app.NewWindow("Hattrick")
	w.SetContent(
		appgui.MakeSplitTab(),
	)
	//widget.NewButton("Quit", func() {app.Quit()})
	w.Resize(fyne.NewSize(1100, 750))
	w.ShowAndRun()
}
