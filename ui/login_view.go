package ui

import (
	"image/color"
	"muscurdig/enums"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetLoginView(state binding.String) *fyne.Container {
	passEntry := widget.NewPasswordEntry()
	errorMsg := canvas.NewText("Wrong password", color.RGBA{255, 0, 0, 255})
	errorMsg.Alignment = fyne.TextAlignCenter
	errorMsg.Hide()

	loginBtn := widget.NewButton("Login", func() {
		entry := passEntry.Text
		passEntry.SetText("")

		if entry == "qwerty" {
			state.Set(enums.List.String())
			return
		}

		errorMsg.Show()
		go func() {
			time.Sleep(time.Millisecond * 800)
			errorMsg.Hide()
		}()
	})

	return container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Insert Master Password"),
			passEntry,
			loginBtn,
			errorMsg,
		),
	)
}
