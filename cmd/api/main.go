package main

import (
	"log"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/app"
)

func main() {
	a := app.New("go-secure-lab")
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
