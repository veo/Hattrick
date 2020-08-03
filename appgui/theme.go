package appgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"image/color"
	"runtime"
)

func font() fyne.Resource {
	if runtime.GOOS == `windows` {
		font, err := fyne.LoadResourceFromPath(`C:\Windows\Fonts\simhei.ttf`)
		if err != nil {
			font = theme.LightTheme().TextFont()
		}
		return font
	} else if runtime.GOOS == `darwin` {
		font, err := fyne.LoadResourceFromPath(`/Library/Fonts/Arial Unicode.ttf`)
		if err != nil {
			font = theme.LightTheme().TextFont()
		}
		return font
	}
	return theme.LightTheme().TextFont()
}

type HatTrickTheme struct{}

func (HatTrickTheme) BackgroundColor() color.Color      { return theme.LightTheme().BackgroundColor() }
func (HatTrickTheme) ButtonColor() color.Color          { return theme.LightTheme().ButtonColor() }
func (HatTrickTheme) HyperlinkColor() color.Color       { return theme.LightTheme().HyperlinkColor() }
func (HatTrickTheme) TextColor() color.Color            { return theme.LightTheme().TextColor() }
func (HatTrickTheme) PlaceHolderColor() color.Color     { return theme.LightTheme().PlaceHolderColor() }
func (HatTrickTheme) PrimaryColor() color.Color         { return theme.LightTheme().PrimaryColor() }
func (HatTrickTheme) FocusColor() color.Color           { return theme.LightTheme().FocusColor() }
func (HatTrickTheme) ScrollBarColor() color.Color       { return theme.LightTheme().ScrollBarColor() }
func (HatTrickTheme) TextSize() int                     { return 11 }
func (HatTrickTheme) TextFont() fyne.Resource           { return font() }
func (HatTrickTheme) TextBoldFont() fyne.Resource       { return font() }
func (HatTrickTheme) TextItalicFont() fyne.Resource     { return font() }
func (HatTrickTheme) TextBoldItalicFont() fyne.Resource { return font() }
func (HatTrickTheme) TextMonospaceFont() fyne.Resource  { return font() }
func (HatTrickTheme) Padding() int                      { return theme.LightTheme().Padding() }
func (HatTrickTheme) IconInlineSize() int               { return theme.LightTheme().IconInlineSize() }
func (HatTrickTheme) ScrollBarSize() int                { return theme.LightTheme().ScrollBarSize() }
func (HatTrickTheme) DisabledButtonColor() color.Color {
	return theme.LightTheme().DisabledButtonColor()
}
func (HatTrickTheme) DisabledIconColor() color.Color { return theme.LightTheme().DisabledIconColor() }
func (HatTrickTheme) DisabledTextColor() color.Color { return theme.LightTheme().DisabledTextColor() }
func (HatTrickTheme) HoverColor() color.Color        { return theme.LightTheme().HoverColor() }
func (HatTrickTheme) IconColor() color.Color         { return theme.LightTheme().IconColor() }
func (HatTrickTheme) ScrollBarSmallSize() int        { return theme.LightTheme().ScrollBarSmallSize() }
func (HatTrickTheme) ShadowColor() color.Color       { return theme.LightTheme().ShadowColor() }
