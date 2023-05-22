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

func GetSettingsView(ctx *c.AppContext) *fyne.Container {
	message := newFlashTxtPlaceholder()
	deleteAllBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		dialog.ShowConfirm("Confirmation", "Are you sure you want to remove all the saved password?",
			func(b bool) {
				if !b {
					return
				}

				ctx.Db.Drop()
				ctx.NavigateTo(c.Setup)
			},
			ctx.GetWindow(),
		)
	})

	exportPassBtn := widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
			path := lu.Path()
			path, err1 := ctx.Db.GenerateDump(path)
			if err1 != nil {
				errorMessage("Could not create dump file!", message)
			} else {
				successMessage(fmt.Sprintf("Saved as: %s", path), message)
			}

		}, ctx.GetWindow())
	})

	importPassBtn := widget.NewButtonWithIcon("", theme.UploadIcon(), func() {
		dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			uc.Close()
			path := uc.URI().Path()
			showPasswordDialog(
				"Insert Dump Master Password", "Password...",
				func(p string) {
					err := ctx.Db.ImportDump(p, path)
					if err != nil {
						errorMessage(fmt.Sprintf("Error: %s", err), message)
					} else {
						successMessage("Dump file imported!", message)
					}

				},
				ctx.GetWindow())

		}, ctx.GetWindow())
	})

	return container.NewMax(
		container.NewBorder(
			centered(h1("Settings")),
			container.NewBorder(
				nil,
				nil,
				backButton(ctx, c.List),
				nil,
				centered(message),
			),
			nil,
			nil,
			container.NewCenter(
				container.NewVBox(
					container.NewGridWithColumns(2,
						widget.NewLabel("Delete all data"),
						deleteAllBtn,
					),
					container.NewGridWithColumns(2,
						widget.NewLabel("Export data"),
						exportPassBtn,
					),
					container.NewGridWithColumns(2,
						widget.NewLabel("Import data"),
						importPassBtn,
					),
				),
			),
		),
	)

}
