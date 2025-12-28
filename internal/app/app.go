package app

import (
	"net/http"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/http/handler"
)

type App struct {
	name string
}

func New(name string) *App {
	return &App{name: name}
}

func (a *App) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Hello)
	return mux
}
