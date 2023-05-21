package app

import (
	"fmt"
	"muscurdig/conf"
	c "muscurdig/context"
	"muscurdig/db"
	"muscurdig/ui"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	isLogEnabled bool
	ctx          *c.AppContext
	views        map[c.AppRoute]func() *fyne.Container
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
	db := db.NewDb(conf.DbFiles)

	ctx := setupContext(db, &w)
	ctx.Version = conf.Version

	a.Settings().SetTheme(&ui.MuscurdigTheme{})

	return App{
		ctx:          &ctx,
		isLogEnabled: isLogEnabled,
		application:  a,
		window:       &w,
		views: map[c.AppRoute]func() *fyne.Container{
			c.Setup:     func() *fyne.Container { return ui.GetSetupView(&ctx) },
			c.Login:     func() *fyne.Container { return ui.GetLoginView(&ctx) },
			c.List:      func() *fyne.Container { return ui.GetPasswordListView(&ctx) },
			c.Details:   func() *fyne.Container { return ui.GetPasswordDetailsView(&ctx) },
			c.AddUpdate: func() *fyne.Container { return ui.GetPasswordAddUpdateView(&ctx) },
			c.Settings:  func() *fyne.Container { return ui.GetSettingsView(&ctx) },
			c.About:     func() *fyne.Container { return ui.GetAboutView(&ctx) },
		},
	}
}

// TODO: for the next project this might be better as a Container
// or Factory with Cache and a stack to simulate pop push routes
func (a *App) getView() *fyne.Container {
	key := a.ctx.CurrentRoute()

	if content, ok := a.views[key]; ok {
		return content()
	}

	return a.views[c.Login]()
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
	a.ctx.OnRouteChange(func() {
		val := a.ctx.CurrentRoute()
		a.log(fmt.Sprintf("route state changed %s", val))
		if val == c.Quit {
			a.application.Quit()
		}

		a.setView()
	})
	a.setView()
	(*a.window).ShowAndRun()

	a.log("exiting...")
}

func (a *App) Cleanup() {
	a.log("Running cleanup")
	a.ctx.Db.Close()
	a.log("cleanup finished")
}

func setupContext(db *db.Db, w *fyne.Window) c.AppContext {
	initialRoute := c.Login
	if _, err := db.GetMasterPassword(); err != nil {
		initialRoute = c.Setup
	}
	return c.NewAppContext(initialRoute, db, w)
}
