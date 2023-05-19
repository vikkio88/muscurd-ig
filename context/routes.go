package context

import (
	"strings"
)

type AppRoute uint8

const (
	login     string = "LOGIN"
	setup     string = "SETUP"
	list      string = "LIST"
	details   string = "DETAILS"
	addUpdate string = "ADDUPDATE"
	about     string = "ABOUT"
	quit      string = "QUIT"

	invalid string = "INVALID_ROUTE"
)

const (
	Login AppRoute = iota
	Setup
	List
	Details
	AddUpdate
	About
	Quit
)

func getMapping() map[AppRoute]string {
	return map[AppRoute]string{
		Login:     login,
		Setup:     setup,
		List:      list,
		Details:   details,
		AddUpdate: addUpdate,
		About:     about,
		Quit:      quit,
	}
}

func getReverseMapping() map[string]AppRoute {
	return map[string]AppRoute{
		login:     Login,
		setup:     Setup,
		list:      List,
		details:   Details,
		addUpdate: AddUpdate,
		about:     About,
		quit:      Quit,
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
