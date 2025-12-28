package main

import (
	"log"
	"net/http"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/app"
)

func main() {
	a := app.New("go-secure-lab")
	serv := http.Server{
		Addr:    ":8088",
		Handler: a.Routes(),
	}
	log.Println("listening on", serv.Addr)

	if err := serv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
