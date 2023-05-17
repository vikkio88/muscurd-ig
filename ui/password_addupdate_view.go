package ui

import (
	"fmt"
	c "muscurdig/context"
	"muscurdig/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetPasswordAddUpdateView(ctx *c.AppContext) *fyne.Container {
	websiteEntry := widget.NewEntry()
	websiteEntry.SetPlaceHolder("Website")
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")
	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("Password")

	isUpdating := false
	headerLabel := "Add New Password"
	var pe models.PasswordEntry
	if id, ok := ctx.RouteParam.(string); ok {
		isUpdating = true
		pe = ctx.Db.GetPasswordById(id)
		headerLabel = "Edit Password"
		websiteEntry.Bind(binding.BindString(&pe.Website))
		usernameEntry.Bind(binding.BindString(&pe.Username))
		passEntry.Bind(binding.BindString(&pe.Password))
	}

	addBtn := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		var err error
		if !isUpdating {
			passEntry := models.NewPasswordEntry(websiteEntry.Text, usernameEntry.Text, passEntry.Text)
			err = ctx.Db.InsertPasswordEntry(passEntry)
		} else {
			err = ctx.Db.UpdatePasswordEntry(pe)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		ctx.NavigateTo(c.List)
	})
	return container.New(layout.NewMaxLayout(),
		container.NewBorder(
			nil,
			container.NewBorder(nil, nil, backButton(ctx, c.List), nil),
			nil, nil,
			container.NewVBox(
				h1(headerLabel),
				websiteEntry,
				usernameEntry,
				passEntry,
				addBtn,
			),
		))
}
