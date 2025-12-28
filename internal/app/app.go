package app

import (
	"context"
	"fmt"
)

type App struct {
	name string
}

func New(name string) *App {
	return &App{name: name}
}
func (a *App) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println(a.name)
		return nil
	}
}
