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

	for i := 0; i < 20; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(50 * time.Millisecond):
			// یک قدم کار (فعلاً فقط هیچ)
		}

	}
	fmt.Println(a.name)
	return nil
}
