package appgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/veo/Hattrick/crypto"
)

func Loadui() {
	app := app.New()
	app.SetIcon(appicon)
	app.Settings().SetTheme(&HatTrickTheme{})
	w := app.NewWindow("Hattrick")
	textbox := &Textbox{}
	//w.SetContent(appgui.Button())
	w.SetContent(widget.NewVBox(
		textbox.InputBox(),
		widget.NewHScrollContainer(fyne.NewContainerWithLayout(layout.NewGridLayout(10),
			widget.NewButton("MD5_Encode_32", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_32(textbox.InputText.Text)) }),
			widget.NewButton("MD5_Encode_16", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_16(textbox.InputText.Text)) }),
			widget.NewButton("SHA1_Encode", func() { textbox.EncodeText.SetText(crypto.SHA1_Encode(textbox.InputText.Text)) }),
			widget.NewButton("SHA256_Encode", func() { textbox.EncodeText.SetText(crypto.SHA256_Encode(textbox.InputText.Text)) }),
			widget.NewButton("SHA512_Encode", func() { textbox.EncodeText.SetText(crypto.SHA512_Encode(textbox.InputText.Text)) }),
			widget.NewButton("16进制(HEX)", func() { textbox.EncodeText.SetText(crypto.HexEncode(textbox.InputText.Text)) }),
			widget.NewButton("Base64_GBK", func() { textbox.EncodeText.SetText(crypto.Base64Encode_GBK(textbox.InputText.Text)) }),
			widget.NewButton("Base64_UTF8", func() { textbox.EncodeText.SetText(crypto.Base64Encode_UTF8(textbox.InputText.Text)) }),
			widget.NewButton("URL编码_GBK", func() { textbox.EncodeText.SetText(crypto.UrlEncode_GBK(textbox.InputText.Text)) }),
			widget.NewButton("URL编码_UTF8", func() { textbox.EncodeText.SetText(crypto.UrlEncode_UTF8(textbox.InputText.Text)) }),
		),
		),
		widget.NewHScrollContainer(fyne.NewContainerWithLayout(layout.NewGridLayout(10),
			widget.NewButton("URL编码_所有", func() { textbox.EncodeText.SetText(crypto.UrlEncode_ALL(textbox.InputText.Text)) }),
			widget.NewButton("ASCII编码", func() { textbox.EncodeText.SetText(crypto.AsciiEncode(textbox.InputText.Text)) }),
			widget.NewButton("ASCII之和", func() { textbox.EncodeText.SetText(crypto.AsciiSUM(textbox.InputText.Text)) }),
			widget.NewButton("SHA256_Encode", func() { textbox.EncodeText.SetText(crypto.SHA256_Encode(textbox.InputText.Text)) }),
			widget.NewButton("SHA512_Encode", func() { textbox.EncodeText.SetText(crypto.SHA512_Encode(textbox.InputText.Text)) }),
			widget.NewButton("16进制(HEX)", func() { textbox.EncodeText.SetText(crypto.HexEncode(textbox.InputText.Text)) }),
			widget.NewButton("Base64_GBK", func() { textbox.EncodeText.SetText(crypto.Base64Encode_GBK(textbox.InputText.Text)) }),
			widget.NewButton("Base64_UTF8", func() { textbox.EncodeText.SetText(crypto.Base64Encode_UTF8(textbox.InputText.Text)) }),
			widget.NewButton("URL编码_GBK", func() { textbox.EncodeText.SetText(crypto.UrlEncode_GBK(textbox.InputText.Text)) }),
			widget.NewButton("URL编码_UTF8", func() { textbox.EncodeText.SetText(crypto.UrlEncode_UTF8(textbox.InputText.Text)) }),
		),
		),
		textbox.EncodeBox(),
		widget.NewHScrollContainer(fyne.NewContainerWithLayout(layout.NewGridLayout(10),
			widget.NewButton("MD5_Encode_32", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_32(textbox.InputText.Text)) }),
			widget.NewButton("MD5_Encode_16", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_16(textbox.InputText.Text)) }),
			widget.NewButton("SHA1_Encode", func() { textbox.EncodeText.SetText(crypto.SHA1_Encode(textbox.InputText.Text)) }),
			widget.NewButton("SHA256_Encode", func() { textbox.EncodeText.SetText(crypto.SHA256_Encode(textbox.InputText.Text)) }),
			widget.NewButton("SHA512_Encode", func() { textbox.EncodeText.SetText(crypto.SHA512_Encode(textbox.InputText.Text)) }),
			widget.NewButton("16进制(HEX)", func() { textbox.EncodeText.SetText(crypto.HexEncode(textbox.InputText.Text)) }),
			widget.NewButton("Base64_GBK", func() { textbox.EncodeText.SetText(crypto.Base64Encode_GBK(textbox.InputText.Text)) }),
			widget.NewButton("Base64_UTF8", func() { textbox.EncodeText.SetText(crypto.Base64Encode_UTF8(textbox.InputText.Text)) }),
			widget.NewButton("MD5_Encode_32", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_32(textbox.InputText.Text)) }),
			widget.NewButton("MD5_Encode_32", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_32(textbox.InputText.Text)) }),
		),
		),
		textbox.DecodeBox(),
	))
	w.Resize(fyne.NewSize(1200, 750))
	w.ShowAndRun()
}
