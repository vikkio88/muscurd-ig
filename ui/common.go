package ui

import (
	"image/color"
	"muscurdig/context"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func h1(text string) *canvas.Text {
	txt := canvas.NewText(text, color.White)
	txt.TextSize = 20
	return txt
}

func h2(text string) *canvas.Text {
	txt := canvas.NewText(text, color.White)
	txt.TextSize = 18
	return txt
}

func backButton(ctx *context.AppContext, route context.AppRoute) *widget.Button {
	return widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		ctx.NavigateTo(route)
	})
}
