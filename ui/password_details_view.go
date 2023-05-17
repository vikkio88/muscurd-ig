package ui

import (
	c "muscurdig/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetPasswordDetailsView(ctx *c.AppContext) *fyne.Container {
	password := ctx.Db.GetPasswordById(ctx.RouteParam.(string))
	passEntry := widget.NewPasswordEntry()
	passEntry.SetText(password.Password)
	passEntry.Disable()
	return container.New(layout.NewMaxLayout(),
		container.NewBorder(
			nil,
			container.NewBorder(nil, nil,
				backButton(ctx, c.List),
				container.NewHBox(
					widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
						ctx.NavigateToWithParam(c.AddUpdate, password.Id)
					}),
					widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
						ctx.Db.DeletePasswordEntry(password.Id)
						ctx.NavigateTo(c.List)
					}),
				),
			),
			nil, nil,
			container.NewCenter(
				container.NewMax(
					container.NewVBox(
						h1("Password Details"),
						h2(password.Website),
						container.NewGridWithColumns(2,
							h2(password.Username),
							widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
								ctx.GetClipboard().SetContent(password.Username)
							}),
						),
						container.NewGridWithColumns(2,
							passEntry,
							widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
								ctx.GetClipboard().SetContent(password.Password)
							}),
						),
					),
				),
			),
		))
}
