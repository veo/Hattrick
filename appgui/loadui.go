package appgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/veo/Hattrick/crypto"
	"net/url"
)

func Loadui() {
	app := app.New()
	//app.SetIcon(fyne.NewStaticResource("Icon.png", static.IconPNG))
	app.Settings().SetTheme(&HatTrickTheme{})
	w := app.NewWindow("Hattrick")
	blog, err := url.Parse("https://xwhoami.com")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	github, err := url.Parse("https://github.com/veo")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	textbox := &Textbox{}
	w.SetContent(widget.NewVBox(
		textbox.InputBox(),
		widget.NewHScrollContainer(fyne.NewContainerWithLayout(layout.NewGridLayout(10),
			widget.NewButton("MD5_Encode_32", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_32(textbox.InputText.Text)) }),
			widget.NewButton("MD5_Encode_16", func() { textbox.EncodeText.SetText(crypto.MD5_Encode_16(textbox.InputText.Text)) }),
			widget.NewButton("SHA1_Encode", func() { textbox.EncodeText.SetText(crypto.SHA1_Encode(textbox.InputText.Text)) }),
			widget.NewButton("SHA256_Encode", func() { textbox.EncodeText.SetText(crypto.SHA256_Encode(textbox.InputText.Text)) }),
			widget.NewButton("SHA512_Encode", func() { textbox.EncodeText.SetText(crypto.SHA512_Encode(textbox.InputText.Text)) }),
			widget.NewButton("Base64_GBK", func() { textbox.EncodeText.SetText(crypto.Base64Encode_GBK(textbox.InputText.Text)) }),
			widget.NewButton("Base64_UTF8", func() { textbox.EncodeText.SetText(crypto.Base64Encode_UTF8(textbox.InputText.Text)) }),
			widget.NewButton("16进制(HEX)", func() { textbox.EncodeText.SetText(crypto.HexEncode(textbox.InputText.Text)) }),
			widget.NewButton("ASCII编码", func() { textbox.EncodeText.SetText(crypto.AsciiEncode(textbox.InputText.Text)) }),
			widget.NewButton("ASCII之和", func() { textbox.EncodeText.SetText(crypto.AsciiSUM(textbox.InputText.Text)) }),
		),
		),
		widget.NewHScrollContainer(fyne.NewContainerWithLayout(layout.NewGridLayout(10),
			widget.NewButton("FromCharCode", func() { textbox.EncodeText.SetText(crypto.FromChar(textbox.InputText.Text)) }),
			widget.NewButton("HTML实体(DEC)", func() { textbox.EncodeText.SetText(crypto.HtmlEncode_Dec(textbox.InputText.Text)) }),
			widget.NewButton("HTML实体(HEX)", func() { textbox.EncodeText.SetText(crypto.HtmlEncode_Hex(textbox.InputText.Text)) }),
			widget.NewButton("Unicode", func() { textbox.EncodeText.SetText(crypto.Unicode(textbox.InputText.Text)) }),
			widget.NewButton("URL编码_所有", func() { textbox.EncodeText.SetText(crypto.UrlEncode_ALL(textbox.InputText.Text)) }),
			widget.NewButton("URL编码_GBK", func() { textbox.EncodeText.SetText(crypto.UrlEncode_GBK(textbox.InputText.Text)) }),
			widget.NewButton("URL编码_UTF8", func() { textbox.EncodeText.SetText(crypto.UrlEncode_UTF8(textbox.InputText.Text)) }),
			widget.NewButton("JsHex", func() { textbox.EncodeText.SetText(crypto.JsHex(textbox.InputText.Text)) }),
			widget.NewButton("JsUnicode", func() { textbox.EncodeText.SetText(crypto.JsUnicode(textbox.InputText.Text)) }),
			widget.NewButton("Unescape", func() { textbox.EncodeText.SetText(crypto.Unescape(textbox.InputText.Text)) }),
		),
		),
		textbox.EncodeBox(),
		widget.NewHScrollContainer(fyne.NewContainerWithLayout(layout.NewGridLayout(10),
			widget.NewButton("Base64_GBK", func() { textbox.DecodeText.SetText(crypto.Base64Decode_GBK(textbox.EncodeText.Text)) }),
			widget.NewButton("Base64_UTF8", func() { textbox.DecodeText.SetText(crypto.Base64Decode_UTF8(textbox.EncodeText.Text)) }),
			widget.NewButton("URL_GBK", func() { textbox.DecodeText.SetText(crypto.UrlDecode_GBK(textbox.EncodeText.Text)) }),
			widget.NewButton("URL_UTF8", func() { textbox.DecodeText.SetText(crypto.UrlDecode_UTF8(textbox.EncodeText.Text)) }),
		),
		),
		textbox.DecodeBox(),
		widget.NewHBox(layout.NewSpacer(), widget.NewLabel("作者: veo "), widget.NewHyperlink("blog", blog), widget.NewHyperlink("github", github), layout.NewSpacer()),
	))
	w.Resize(fyne.NewSize(1200, 750))
	w.ShowAndRun()
}
