package main

import (
	"chatservice/internal/app"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	app := app.NewApp(ctx)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
