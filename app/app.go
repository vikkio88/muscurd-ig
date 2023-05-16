package app

import (
	"fmt"
	"muscurdig/conf"
	"muscurdig/state"
	s "muscurdig/state"
	"muscurdig/ui"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	isLogEnabled bool
	state        *s.AppState
	views        map[s.AppRoute]*fyne.Container
	application  fyne.App
	window       *fyne.Window
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

	return App{
		state:        &initialState,
		isLogEnabled: isLogEnabled,
		application:  a,
		window:       &w,
		views: map[s.AppRoute]*fyne.Container{
			s.Setup:   ui.GetSetupView(&initialState),
			s.Login:   ui.GetLoginView(&initialState),
			s.List:    ui.GetPasswordListView(&initialState),
			s.Details: ui.GetPasswordDetailsView(&initialState),
		},
	}
}

// TODO: for the next project this might be better as a Container
// or Factory with Cache and a stack to simulate pop push routes
func (a *App) getView() *fyne.Container {
	key := a.state.CurrentRoute()

	if content, ok := a.views[key]; ok {
		return content
	}

	return a.views[s.Login]
}
func (a *App) setView() {
	(*a.window).SetContent(a.getView())
}

func (a *App) log(msg string) {
	if a.isLogEnabled {
		fmt.Printf("%s - %s\n", time.Now().Format("15:04:05"), msg)
	}
}

func (a *App) Run() {
	a.state.OnRouteChange(func() {
		val := a.state.CurrentRoute()
		a.log(fmt.Sprintf("route state changed %s", val))
		if val == state.Quit {
			a.application.Quit()
		}

		a.setView()
	})
	a.setView()
	(*a.window).ShowAndRun()
}
