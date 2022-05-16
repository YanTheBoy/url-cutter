package main

import (
	"github.com/iliarkhpv/url-cutter/internal/app"
	config "github.com/iliarkhpv/url-cutter/internal/cfg"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Server is not running due to: ", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatal("Some problems with server running", err)
	}
}
