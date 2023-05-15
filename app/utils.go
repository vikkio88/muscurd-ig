package app

import (
	"muscurdig/enums"
	"os"
)

func getInitialState() string {
	initialState := enums.Login.String()
	if _, err := os.Stat("muscurdi_db"); os.IsNotExist(err) {
		initialState = enums.Setup.String()
	}

	return initialState
}
