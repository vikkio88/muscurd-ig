package context

import (
	"muscurdig/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type AppContext struct {
	Route      binding.String
	RouteParam any
	Db         *db.Db
	w          *fyne.Window
}

func NewAppContext(initialRoute AppRoute, db *db.Db, window *fyne.Window) AppContext {
	route := initialRoute.String()
	return AppContext{
		Route: binding.BindString(&route),
		Db:    db,
		w:     window,
	}
}
func (s *AppContext) GetClipboard() fyne.Clipboard {
	w := s.w
	return (*w).Clipboard()
}

// You could add multiple OnrouteChange
// todo: maybe needs to find a better name for this
func (s *AppContext) OnRouteChange(callback func()) {
	s.Route.AddListener(binding.NewDataListener(callback))
}

func (s *AppContext) CurrentRoute() AppRoute {
	r, _ := s.Route.Get()
	return RouteFromString(r)
}

func (s *AppContext) NavigateTo(route AppRoute) {
	s.RouteParam = nil
	s.Route.Set(route.String())
}

func (s *AppContext) NavigateToWithParam(route AppRoute, param any) {
	s.RouteParam = param
	s.Route.Set(route.String())
}
