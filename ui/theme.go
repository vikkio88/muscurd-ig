package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MuscurdigTheme struct {
}

func (m *MuscurdigTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameDisabled:
		if v == theme.VariantLight {
			return &color.RGBA{R: 33, G: 33, B: 33, A: 255}
		}
		return &color.RGBA{R: 255, G: 255, B: 255, A: 255}
	}

	return theme.DefaultTheme().Color(n, v)
}

func (m *MuscurdigTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTextFont()
}

func (m *MuscurdigTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *MuscurdigTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
