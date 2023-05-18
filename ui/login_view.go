package ui

import (
	c "muscurdig/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetLoginView(ctx *c.AppContext) *fyne.Container {
	passEntry := widget.NewPasswordEntry()
	errorMsg := newFlashTxtPlaceholder()
	errorMsg.Alignment = fyne.TextAlignCenter

	loginBtn := widget.NewButton("Login", func() {
		entry := passEntry.Text
		passEntry.SetText("")

		masterPassword, err := ctx.Db.GetMasterPassword()
		if err != nil {
			panic(err)
		}

		if masterPassword.Check(entry) {
			ctx.NavigateTo(c.List)
			return
		}
		errorMessage("Wrong password", errorMsg)
	})

	return container.New(layout.NewCenterLayout(),
		container.NewVBox(
			widget.NewLabel("Insert Master Password"),
			passEntry,
			loginBtn,
			errorMsg,
		),
	)
}
