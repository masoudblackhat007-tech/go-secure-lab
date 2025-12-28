package app

import (
	"context"
	"fmt"
	"time"
)

type App struct {
	name string
}

func New(name string) *App {
	return &App{name: name}
}
func (a *App) Run(ctx context.Context) error {
	time.Sleep(500 * time.Millisecond)
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println(a.name)
		return nil
	}
}
