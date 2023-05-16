package state

import (
	"strings"
)

type AppRoute uint8

const (
	login   string = "LOGIN"
	setup   string = "SETUP"
	list    string = "LIST"
	details string = "DETAILS"
	quit    string = "QUIT"

	invalid string = "INVALID_ROUTE"
)

const (
	Login AppRoute = iota
	Setup
	List
	Details
	Quit
)

func getMapping() map[AppRoute]string {
	return map[AppRoute]string{
		Login:   login,
		Setup:   setup,
		List:    list,
		Details: details,
		Quit:    quit,
	}
}

func getReverseMapping() map[string]AppRoute {
	return map[string]AppRoute{
		login:   Login,
		setup:   Setup,
		list:    List,
		details: Details,
		quit:    Quit,
	}
}

func RouteFromString(route string) AppRoute {
	route = strings.ToUpper(route)
	mapping := getReverseMapping()
	if val, ok := mapping[route]; ok {
		return val
	}

	return Login
}

func (a AppRoute) String() string {
	mapping := getMapping()
	if val, ok := mapping[a]; ok {
		return val
	}

	return invalid
}
