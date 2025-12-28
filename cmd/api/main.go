package main

import (
	"context"
	"log"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/app"
)

func main() {
	ctx := context.Background()

	a := app.New("go-secure-lab")
	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
