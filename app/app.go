package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type appState uint8

const (
	login appState = iota
	view
	quit
)

func (a appState) String() string {
	switch a {
	case login:
		return "LOGIN"
	case view:
		return "VIEW"
	case quit:
		return "QUIT"
	}

	return "INVALID_STATE"
}

type App struct {
	state appState
}

func NewApp() App {
	return App{
		state: login,
	}
}

func (a *App) Run() {
	fmt.Println("[q] to quit")
	for a.state != quit {
		fmt.Printf("state: %s\n\n", a.state)
		c := getStr("> ")

		if c == "q" {
			a.state = quit
		}
	}
}

func getStr(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
