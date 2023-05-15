package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetLoginView(state binding.String) *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Login View"),
		getNavigation(state),
	)
}
