package main

import (
	"context"
	"log"
	"time"

	"github.com/masoudblackhat007-tech/go-secure-lab/internal/app"
)

func main() {
	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	a := app.New("go-secure-lab")
	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
