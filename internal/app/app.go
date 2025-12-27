package app

import "fmt"

type App struct {
	name string
}

func New(name string) *App {
	return &App{name: name}
}

func (a *App) Run() error {
	fmt.Println(a.name)
	return nil
}
