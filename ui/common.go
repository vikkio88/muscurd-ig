package ui

import (
	"image/color"
	"muscurdig/context"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

func small(text string) *canvas.Text {
	txt := canvas.NewText(text, color.White)
	txt.TextSize = 10
	return txt
}

func errorMessage(errorMsg string, textItem *canvas.Text) {
	flashMessage(
		errorMsg, textItem,
		time.Millisecond*800,
		color.RGBA{255, 0, 50, 255},
	)

}
func successMessage(msg string, textItem *canvas.Text) {
	flashMessage(
		msg, textItem,
		time.Millisecond*800,
		color.RGBA{0, 255, 50, 255},
	)

}
func flashMessage(msg string, textItem *canvas.Text, duration time.Duration, color color.Color) {
	textItem.Color = color
	textItem.Text = msg
	go func() {
		time.Sleep(duration)
		textItem.Text = ""
	}()
}

func newFlashTxtPlaceholder() *canvas.Text {
	return canvas.NewText("", color.White)
}

func backButton(ctx *context.AppContext, route context.AppRoute) *widget.Button {
	return widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		ctx.NavigateTo(route)
	})
}

func centered(object fyne.CanvasObject) *fyne.Container {
	return container.NewCenter(object)
}
func rightAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, nil, object)
}

func leftAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, object, nil)
}

func topAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(object, nil, nil, nil)
}

func bottomAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, object, nil, nil)
}
