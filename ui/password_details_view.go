package ui

import (
	s "muscurdig/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetPasswordDetailsView(state *s.AppState) *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Password Details View"),
	)
}
