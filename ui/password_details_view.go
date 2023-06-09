package ui

import (
	"fmt"
	c "muscurdig/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetPasswordDetailsView(ctx *c.AppContext) *fyne.Container {
	successMsg := newFlashTxtPlaceholder()
	password := ctx.Db.GetPasswordById(ctx.RouteParam.(string))
	passEntry := widget.NewPasswordEntry()
	passEntry.SetText(password.Password)
	passEntry.Disable()

	copyUsernameBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		ctx.GetClipboard().SetContent(password.Username)
		successMessage("Username copied to clipboard.", successMsg)
	})
	copyPasswordBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		ctx.GetClipboard().SetContent(password.Password)
		successMessage("Password copied to clipboard.", successMsg)
	})

	return container.NewMax(
		container.NewBorder(
			nil,
			container.NewBorder(nil, nil,
				backButton(ctx, c.List),
				container.NewHBox(
					widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
						ctx.NavigateToWithParam(c.AddUpdate, password.Id)
					}),
					widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
						dialog.ShowConfirm(
							fmt.Sprintf("Deleting password for \"%s\"", password.Website),
							"Are you sure?",
							func(b bool) {
								if b {
									ctx.Db.DeletePasswordEntry(password.Id)
									ctx.NavigateTo(c.List)
								}
							},
							ctx.GetWindow(),
						)
					}),
				),
				container.NewCenter(successMsg),
			),
			nil, nil,
			container.NewMax(
				container.NewVBox(
					container.NewCenter(h1("Password Details")),
					h2(password.Website),
					container.NewBorder(
						nil,
						nil,
						nil,
						copyUsernameBtn,
						h2(password.Username),
					),
					container.NewBorder(
						nil,
						nil,
						nil,
						copyPasswordBtn,
						passEntry,
					),
				),
			),
		))
}
