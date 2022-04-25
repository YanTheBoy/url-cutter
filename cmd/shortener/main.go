package main

import (
	"github.com/iliarkhpv/url-cutter/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal("Some problems with server running", err)
	}
}
