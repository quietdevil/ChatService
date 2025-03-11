package main

import (
	"context"
	"github.com/quietdevil/ChatSevice/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
