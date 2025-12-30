package app

import (
	"net/http"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/http/handler"
	"github.com/masoudblackhat007-tech/go-secure-lab/internal/http/middleware"
)

type App struct {
	name string
}

func New(name string) *App {
	return &App{name: name}
}

func (a *App) Routes() http.Handler {
	mux := http.NewServeMux()
	h := &handler.Handler{AppName: a.name}
	mux.HandleFunc("/", h.Hello)
	mux.HandleFunc("/health", h.Health)

	// اعمال Middleware روی کل مسیرها
	return middleware.Logging(mux)
}
