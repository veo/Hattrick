package appgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"image/color"
)

type HatTrickTheme struct{}

func (HatTrickTheme) TextFont() fyne.Resource     { return Font }
func (HatTrickTheme) TextBoldFont() fyne.Resource { return Font }

func (HatTrickTheme) BackgroundColor() color.Color { return theme.LightTheme().BackgroundColor() }
func (HatTrickTheme) ButtonColor() color.Color     { return theme.LightTheme().ButtonColor() }
func (HatTrickTheme) DisabledButtonColor() color.Color {
	return theme.LightTheme().DisabledButtonColor()
}
func (HatTrickTheme) IconColor() color.Color         { return theme.LightTheme().IconColor() }
func (HatTrickTheme) DisabledIconColor() color.Color { return theme.LightTheme().DisabledIconColor() }
func (HatTrickTheme) HyperlinkColor() color.Color    { return theme.LightTheme().HyperlinkColor() }
func (HatTrickTheme) TextColor() color.Color         { return theme.LightTheme().TextColor() }
func (HatTrickTheme) DisabledTextColor() color.Color { return theme.LightTheme().DisabledTextColor() }
func (HatTrickTheme) HoverColor() color.Color        { return theme.LightTheme().HoverColor() }
func (HatTrickTheme) PlaceHolderColor() color.Color  { return theme.LightTheme().PlaceHolderColor() }
func (HatTrickTheme) PrimaryColor() color.Color      { return theme.LightTheme().PrimaryColor() }
func (HatTrickTheme) FocusColor() color.Color        { return theme.LightTheme().FocusColor() }
func (HatTrickTheme) ScrollBarColor() color.Color    { return theme.LightTheme().ScrollBarColor() }
func (HatTrickTheme) ShadowColor() color.Color       { return theme.LightTheme().ShadowColor() }
func (HatTrickTheme) TextSize() int                  { return theme.LightTheme().TextSize() }
func (HatTrickTheme) TextItalicFont() fyne.Resource  { return theme.LightTheme().TextItalicFont() }
func (HatTrickTheme) TextBoldItalicFont() fyne.Resource {
	return theme.LightTheme().TextBoldItalicFont()
}
func (HatTrickTheme) TextMonospaceFont() fyne.Resource { return theme.LightTheme().TextMonospaceFont() }
func (HatTrickTheme) Padding() int                     { return theme.LightTheme().Padding() }
func (HatTrickTheme) IconInlineSize() int              { return theme.LightTheme().IconInlineSize() }
func (HatTrickTheme) ScrollBarSize() int               { return theme.LightTheme().ScrollBarSize() }
func (HatTrickTheme) ScrollBarSmallSize() int          { return theme.LightTheme().ScrollBarSmallSize() }
