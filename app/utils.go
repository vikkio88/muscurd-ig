package app

import (
	"muscurdig/state"
	"os"
)

func getInitialState() state.AppState {
	initialRoute := state.Login
	if _, err := os.Stat("muscurdi_db"); os.IsNotExist(err) {
		initialRoute = state.Setup
	}
	return state.NewAppState(initialRoute)
}
