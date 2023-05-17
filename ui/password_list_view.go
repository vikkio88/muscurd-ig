package ui

import (
	"fmt"
	c "muscurdig/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetPasswordListView(ctx *c.AppContext) *fyne.Container {
	passwordCount := ctx.Db.GetPasswordCount()
	searchEntry := widget.NewEntry()
	content := container.New(layout.NewMaxLayout(), container.NewCenter(widget.NewLabel(fmt.Sprintf("%d Password Entries", passwordCount))))
	searchBtn := widget.NewButtonWithIcon("Search", theme.SearchIcon(), func() {
		passwords := ctx.Db.GetAllPasswords()
		content.RemoveAll()
		content.Add(
			// TODO: need to do this a Bindable List so removing items will refresh here too
			widget.NewList(
				func() int { return len(passwords) },
				getPasswordEntryRow,
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					label, btnView, btnEdit, btnDelete := getListWidgetsFromContainer(&co)
					pe := passwords[lii]
					label.SetText(pe.Website)

					btnView.OnTapped = func() {
						ctx.NavigateToWithParam(c.Details, pe.Id)
					}

					btnEdit.OnTapped = func() {
						ctx.NavigateToWithParam(c.AddUpdate, pe.Id)
					}
					btnDelete.OnTapped = func() {
						ctx.Db.DeletePasswordEntry(pe.Id)
						ctx.NavigateTo(c.List)
					}
				},
			),
		)
	})

	return container.New(layout.NewMaxLayout(),
		container.NewBorder(
			container.NewGridWithColumns(2, searchEntry, searchBtn),
			container.NewBorder(nil, nil, nil, widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
				ctx.NavigateTo(c.AddUpdate)
			})),
			nil, nil,
			content,
		),
	)
}

func getPasswordEntryRow() fyne.CanvasObject {
	return container.NewGridWithColumns(3,
		widget.NewLabel(""),
		container.NewMax(),
		container.NewBorder(
			nil,
			nil,
			nil,
			container.NewHBox(
				widget.NewButtonWithIcon("", theme.FileTextIcon(), func() {}),
				widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {}),
				widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
			),
		),
	)
}
func getListWidgetsFromContainer(co *fyne.CanvasObject) (*widget.Label, *widget.Button, *widget.Button, *widget.Button) {
	container := (*co).(*fyne.Container)
	label := container.Objects[0].(*widget.Label)
	btnBoxWrapper := container.Objects[2].(*fyne.Container)
	btnBox := btnBoxWrapper.Objects[0].(*fyne.Container)
	btnView := btnBox.Objects[0].(*widget.Button)
	btnEdit := btnBox.Objects[1].(*widget.Button)
	btnDelete := btnBox.Objects[2].(*widget.Button)

	return label, btnView, btnEdit, btnDelete
}
