package main

import (
	"context"
	"log"
	"wb/internal/app"
)

func main() {
	ctx := context.Background()
	application := app.Create()
	err := application.Start(ctx)

	if err != nil {
		log.Println(err)
		return
	}
}
