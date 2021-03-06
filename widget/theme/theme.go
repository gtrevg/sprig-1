package theme

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/widget/material"
)

var (
	teal         = color.RGBA{R: 0x44, G: 0xa8, B: 0xad, A: 255}
	brightTeal   = color.RGBA{R: 0x79, G: 0xda, B: 0xdf, A: 255}
	darkTeal     = color.RGBA{R: 0x00, G: 0x79, B: 0x7e, A: 255}
	green        = color.RGBA{R: 0x45, G: 0xae, B: 0x7f, A: 255}
	brightGreen  = color.RGBA{R: 0x79, G: 0xe0, B: 0xae, A: 255}
	darkGreen    = color.RGBA{R: 0x00, G: 0x7e, B: 0x52, A: 255}
	gold         = color.RGBA{R: 255, G: 214, B: 79, A: 255}
	lightGold    = color.RGBA{R: 255, G: 255, B: 129, A: 255}
	darkGold     = color.RGBA{R: 200, G: 165, B: 21, A: 255}
	white        = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	lightGray    = color.RGBA{R: 225, G: 225, B: 225, A: 255}
	darkGray     = color.RGBA{R: 100, G: 100, B: 100, A: 255}
	veryDarkGray = color.RGBA{R: 50, G: 50, B: 50, A: 255}
	black        = color.RGBA{A: 255}

	purple1           = color.RGBA{R: 69, G: 56, B: 127, A: 255}
	lightPurple1      = color.RGBA{R: 121, G: 121, B: 174, A: 255}
	darkPurple1       = color.RGBA{R: 99, G: 41, B: 115, A: 255}
	purple2           = color.RGBA{R: 127, G: 96, B: 183, A: 255}
	lightPurple2      = color.RGBA{R: 121, G: 150, B: 223, A: 255}
	darkPurple2       = color.RGBA{R: 101, G: 89, B: 223, A: 255}
	dmBackground      = color.RGBA{R: 12, G: 12, B: 15, A: 255}
	dmDarkBackground  = black
	dmLightBackground = color.RGBA{R: 27, G: 22, B: 33, A: 255}
	dmText            = color.RGBA{R: 194, G: 196, B: 199, A: 255}
)

func New() *Theme {
	gioTheme := material.NewTheme(gofont.Collection())
	var t Theme
	t.Theme = gioTheme
	t.Primary = Colors{
		Default: green,
		Light:   brightGreen,
		Dark:    darkGreen,
	}
	t.Secondary = Colors{
		Default: teal,
		Light:   brightTeal,
		Dark:    darkTeal,
	}
	t.Background = Colors{
		Default: lightGray,
		Light:   white,
		Dark:    black,
	}
	t.Theme.Color.Primary = t.Primary.Default
	t.Ancestors = &t.Secondary.Default
	t.Descendants = &t.Secondary.Default
	t.Selected = &t.Secondary.Light
	t.Unselected = &t.Background.Light
	t.Siblings = t.Unselected
	return &t
}

func (t *Theme) ToDark() {
	t.Background.Dark = darkGray
	t.Background.Default = veryDarkGray
	t.Background.Light = black
	t.Color.Text = white
	t.Color.InvText = black
	t.Color.Hint = lightGray
	t.Primary.Default = purple1
	t.Primary.Light = lightPurple1
	t.Primary.Dark = darkPurple1
	t.Theme.Color.Primary = t.Primary.Default
	t.Secondary.Default = purple2
	t.Secondary.Light = lightPurple2
	t.Secondary.Dark = darkPurple2

	t.Background.Default = dmBackground
	t.Background.Light = dmLightBackground
	t.Background.Dark = dmDarkBackground

	t.Theme.Color.Text = dmText
	t.Theme.Color.InvText = white
}

type Theme struct {
	*material.Theme
	Primary    Colors
	Secondary  Colors
	Background Colors

	Ancestors, Descendants, Selected, Siblings, Unselected *color.RGBA
}

type Colors struct {
	Default, Light, Dark color.RGBA
}
