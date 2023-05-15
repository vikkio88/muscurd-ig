package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetPasswordListView(state binding.String) *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Password List View"),
		getNavigation(state),
	)
}
