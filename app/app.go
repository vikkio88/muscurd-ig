package app

import (
	"fmt"
	"muscurdig/enums"
	"muscurdig/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

type App struct {
	state       binding.String
	views       map[string]*fyne.Container
	application fyne.App
	window      *fyne.Window
	content     *fyne.Container
}

func NewApp() App {
	// TODO: Move this to env or config
	a := app.NewWithID("muscurd-ig_main")
	w := a.NewWindow("Muscurdi - micro password manager")
	w.Resize(fyne.NewSize(450, 300))
	w.SetFixedSize(true)
	// all of this should be configurable

	initialState := enums.Login.String()
	state := binding.BindString(&initialState)
	return App{
		state:       state,
		application: a,
		window:      &w,
		views: map[string]*fyne.Container{
			enums.Login.String():   ui.GetLoginView(state),
			enums.List.String():    ui.GetPasswordListView(state),
			enums.Details.String(): ui.GetPasswordDetailsView(state),
		},
	}
}

func (a *App) GetView() *fyne.Container {
	key, _ := a.state.Get()

	if content, ok := a.views[key]; ok {
		return content
	}

	return a.views[enums.Login.String()]
}

func (a *App) Run() {
	a.state.AddListener(binding.NewDataListener(func() {
		// change this to be a structbind
		val, _ := a.state.Get()
		fmt.Printf("changed: %s\n", val)
		if val == enums.Quit.String() {
			a.application.Quit()
		}

		(*a.window).SetContent(a.GetView())
	}))
	w := *a.window
	a.content = a.GetView()
	w.SetContent(a.content)
	w.ShowAndRun()
}
