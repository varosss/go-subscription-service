package main

import (
	"go-subscription-service/internal/infrastructure/builder"
	"go-subscription-service/internal/infrastructure/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app, err := builder.BuildApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
