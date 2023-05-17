package ui

import (
	"muscurdig/conf"
	c "muscurdig/context"
	"muscurdig/libs"
	"muscurdig/models"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetSetupView(ctx *c.AppContext) *fyne.Container {
	passEntry := widget.NewPasswordEntry()

	loginBtn := widget.NewButton("Login", func() {
		entry := passEntry.Text
		os.Mkdir(conf.DbFiles, 0700)

		crypto := libs.NewCrypto(entry)
		crypted, errCrypt := crypto.EncryptB64(entry)
		if errCrypt != nil {
			panic(errCrypt)
		}
		_, err := ctx.Db.SaveMasterPassword(models.NewMasterPasswordFromB64(crypted))
		if err != nil {
			panic(err)
		}
		ctx.NavigateTo(c.List)
	})

	return container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Generate Master Password"),
			passEntry,
			loginBtn,
		),
	)
}
