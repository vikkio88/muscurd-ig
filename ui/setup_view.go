package ui

import (
	"fmt"
	"muscurdig/conf"
	"muscurdig/libs"
	s "muscurdig/state"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetSetupView(state *s.AppState) *fyne.Container {
	passEntry := widget.NewPasswordEntry()

	loginBtn := widget.NewButton("Login", func() {
		entry := passEntry.Text
		os.Mkdir(conf.DbFiles, 0700)

		//TODO: This is temporary until I set clover
		crypto := libs.NewCrypto(entry)
		crypted, errCrypt := crypto.EncryptB64(entry)
		if errCrypt != nil {
			panic(errCrypt)
		}
		err := os.WriteFile(fmt.Sprintf("./%s/pwd", conf.DbFiles), []byte(crypted), 0644)
		if err != nil {
			panic(err)
		}
		state.NavigateTo(s.List)
	})

	return container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Generate Master Password"),
			passEntry,
			loginBtn,
		),
	)
}
