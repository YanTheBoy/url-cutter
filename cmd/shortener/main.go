package main

import (
	"flag"
	"github.com/iliarkhpv/url-cutter/internal/app"
	config "github.com/iliarkhpv/url-cutter/internal/cfg"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Server is not running due to: ", err)
	}

	parseFlags(cfg)

	if err := app.Run(cfg); err != nil {
		log.Fatal("Some problems with server running", err)
	}
}

func parseFlags(cfg *config.Config) {
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "SERVER_ADDRESS")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "BASE_URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "FILE_STORAGE_PATH")
	flag.Parse()
}
