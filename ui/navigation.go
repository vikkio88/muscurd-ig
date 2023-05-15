package ui

import (
	"muscurdig/enums"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func getNavigation(state binding.String) *fyne.Container {

	toList := widget.NewButton("To List", func() {
		state.Set(enums.List.String())
	})

	toLogin := widget.NewButton("To Login", func() {
		state.Set(enums.Login.String())
	})

	toDetails := widget.NewButton("To Details", func() {
		state.Set(enums.Details.String())
	})

	quit := widget.NewButton("Exit", func() {
		state.Set(enums.Quit.String())
	})
	state.AddListener(binding.NewDataListener(func() {
		current, _ := state.Get()
		if current == enums.List.String() {
			toList.Disable()
		} else {
			toList.Enable()
		}

		if current == enums.Login.String() {
			toLogin.Disable()
		} else {
			toLogin.Enable()
		}
		if current != enums.List.String() {
			toDetails.Disable()
		} else {
			toDetails.Enable()
		}

	}))

	return container.New(layout.NewHBoxLayout(),
		toList,
		toDetails,
		toLogin,
		quit,
	)
}
