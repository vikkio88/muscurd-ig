package ui

import (
	"fmt"
	c "muscurdig/context"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetAboutView(ctx *c.AppContext) *fyne.Container {
	url, _ := url.Parse("https://github.com/vikkio88/muscurd-ig")

	return container.NewMax(
		container.NewBorder(
			nil,
			leftAligned(
				backButton(ctx, c.List),
			),
			nil,
			nil,
			container.NewCenter(
				container.NewVBox(
					centered(h1("About")),
					widget.NewLabel("Muscurd-ig is a micro password manager."),
					centered(widget.NewLabel(fmt.Sprintf("version: %s", ctx.Version))),
					container.NewMax(),
					centered(widget.NewHyperlink("github.com/vikkio88/muscurd-ig", url)),
				),
			),
		),
	)

}
