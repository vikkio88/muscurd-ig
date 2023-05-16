package ui

import (
	"fmt"
	s "muscurdig/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetPasswordListView(state *s.AppState) *fyne.Container {
	searchEntry := widget.NewEntry()
	content := container.NewCenter(widget.NewLabel("0 Password Entries"))
	searchBtn := widget.NewButton("Search", func() {
		fmt.Println("Searched")
		content.RemoveAll()
		content.Add(container.NewCenter(widget.NewLabel("1 Password Entries")))
	})

	return container.New(layout.NewMaxLayout(),
		container.NewBorder(
			container.NewGridWithColumns(2, searchEntry, searchBtn),
			nil, nil, nil,
			content,
		),
	)
}
