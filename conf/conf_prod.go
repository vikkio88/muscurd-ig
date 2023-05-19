//go:build prod

package conf

const (
	AppId            = "muscurd-ig_main"
	WindowTitle      = "Muscurdi - micro password manager"
	WindowWidth      = 450
	WindowHeight     = 300
	WindowFixed      = true
	DbFiles          = "muscurdi_db"
	EnableConsoleLog = false
	Version          = "PROD_VERSION"
)

/*
need to add this on settings.json as gopls fails
but if I add it go.mod triggers errors
    "gopls": {
        "build.buildFlags": [
            "prod"
        ]
    }
*/
