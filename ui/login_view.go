package ui

import (
	"image/color"
	"muscurdig/db"
	s "muscurdig/state"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetLoginView(state *s.AppState) *fyne.Container {
	passEntry := widget.NewPasswordEntry()
	errorMsg := canvas.NewText("Wrong password", color.RGBA{255, 0, 0, 255})
	errorMsg.Alignment = fyne.TextAlignCenter
	errorMsg.Hide()

	loginBtn := widget.NewButton("Login", func() {
		entry := passEntry.Text
		passEntry.SetText("")
		masterPassword := db.GetMasterPassword()
		if masterPassword.Check(entry) {
			state.NavigateTo(s.List)
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
