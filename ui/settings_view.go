package ui

import (
	c "muscurdig/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetSettingsView(ctx *c.AppContext) *fyne.Container {
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
		dialog.ShowInformation("Coming soon!", "This featute is coming soon", ctx.GetWindow())
	})

	importPassBtn := widget.NewButtonWithIcon("", theme.UploadIcon(), func() {
		dialog.ShowInformation("Coming soon!", "This featute is coming soon", ctx.GetWindow())
	})

	return container.NewMax(
		container.NewBorder(
			centered(h1("Settings")),
			leftAligned(
				backButton(ctx, c.List),
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
