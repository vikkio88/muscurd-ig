package state

import "fyne.io/fyne/v2/data/binding"

type AppState struct {
	Route binding.String
}

func NewAppState(initialRoute AppRoute) AppState {
	route := initialRoute.String()
	return AppState{
		Route: binding.BindString(&route),
	}
}

// You could add multiple OnrouteChange
// todo: maybe needs to find a better name for this
func (s *AppState) OnRouteChange(callback func()) {
	s.Route.AddListener(binding.NewDataListener(callback))
}

func (s *AppState) CurrentRoute() AppRoute {
	r, _ := s.Route.Get()
	return RouteFromString(r)
}

func (s *AppState) NavigateTo(route AppRoute) {
	s.Route.Set(route.String())
}
