package app

import (
	"fmt"
	"muscurdig/conf"
	"muscurdig/enums"
	"muscurdig/ui"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

type App struct {
	isLogEnabled bool
	state        binding.String
	views        map[string]*fyne.Container
	application  fyne.App
	window       *fyne.Window
	content      *fyne.Container
}

func NewApp() App {
	a := app.NewWithID(conf.AppId)
	w := a.NewWindow(conf.WindowTitle)
	w.Resize(fyne.NewSize(
		conf.WindowWidth,
		conf.WindowHeight,
	))
	w.SetFixedSize(conf.WindowFixed)
	isLogEnabled := conf.EnableConsoleLog

	initialState := getInitialState()
	state := binding.BindString(&initialState)
	return App{
		state:        state,
		isLogEnabled: isLogEnabled,
		application:  a,
		window:       &w,
		views: map[string]*fyne.Container{
			enums.Setup.String():   ui.GetSetupView(state),
			enums.Login.String():   ui.GetLoginView(state),
			enums.List.String():    ui.GetPasswordListView(state),
			enums.Details.String(): ui.GetPasswordDetailsView(state),
		},
	}
}

// TODO: for the next project this might be better as a Container
// or Factory with Cache and a stack to simulate pop push routes
func (a *App) GetView() *fyne.Container {
	key, _ := a.state.Get()

	if content, ok := a.views[key]; ok {
		return content
	}

	return a.views[enums.Login.String()]
}

func (a *App) log(msg string) {
	if a.isLogEnabled {
		fmt.Printf("%s - %s\n", time.Now().Format("15:04:05"), msg)
	}
}

func (a *App) Run() {
	a.state.AddListener(binding.NewDataListener(func() {
		// change this to be a structbind
		val, _ := a.state.Get()
		a.log(fmt.Sprintf("state changed %s", val))
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
