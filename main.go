package main

import "github.com/veo/Hattrick/appgui"

//func main() {
//	app := app.New()
//	app.SetIcon(appgui.Icon)
//	app.Settings().SetTheme(&appgui.HatTrickTheme{})
//	w := app.NewWindow("Hattrick")
//	textbox := &appgui.Textbox{}
//	//w.SetContent(appgui.Button())
//	w.SetContent(widget.NewVBox(
//		textbox.InputBox(),
//		widget.NewButton("Encode", func() {
//			textbox.EncodeText.SetText(textbox.InputText.Text)
//		}),
//		textbox.EncodeBox(),
//		widget.NewButton("Decode", func() {
//			textbox.DecodeText.SetText(textbox.EncodeText.Text)
//		}),
//		textbox.DecodeBox(),
//	))
//	w.Resize(fyne.NewSize(1100, 750))
//	w.ShowAndRun()
//}

func main() {
	appgui.Loadui()
}
