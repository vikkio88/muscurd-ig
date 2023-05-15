package ui

import (
	"fmt"
	"muscurdig/conf"
	"muscurdig/enums"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetSetupView(state binding.String) *fyne.Container {
	passEntry := widget.NewPasswordEntry()

	loginBtn := widget.NewButton("Login", func() {
		entry := passEntry.Text
		os.Mkdir(conf.DbFiles, 0644)

		//TODO: Error this does not seem to work
		os.WriteFile(fmt.Sprintf("%s/pwd", conf.DbFiles), []byte(entry), 0644)
		state.Set(enums.List.String())
	})

	return container.New(layout.NewCenterLayout(),
		container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Generate Master Password"),
			passEntry,
			loginBtn,
		),
	)
}
