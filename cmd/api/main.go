package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("invalid port %q: must be a number", port)
	}
	addr := "127.0.0.1:" + port
	a := app.New("go-secure-lab")
	serv := http.Server{
		Addr:    addr,
		Handler: a.Routes(),
	}
	log.Println("listening on", serv.Addr)

	if err := serv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
