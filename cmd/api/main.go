package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/app"
	"github.com/masoudblackhat007-tech/go-secure-lab/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	port, err := config.GetPort()
	if err != nil {
		log.Fatalf("invalid port configuration: %v", err)
	}

	addr := "127.0.0.1:" + port
	a := app.New("go-secure-lab")
	serv := http.Server{
		Addr:    addr,
		Handler: a.Routes(),
	}
	ln, err := net.Listen("tcp", serv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("Listening on " + ln.Addr().String())
		if err := serv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s\n", err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down the gracefully ...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
