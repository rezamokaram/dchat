package main

import (
	"flag"
	"log"

	grpc "github.com/RezaMokaram/chapp/api/handler/grpc/presence"
	"github.com/RezaMokaram/chapp/app/presence"
	"github.com/RezaMokaram/chapp/config"
)

func main() {
	var path string
	flag.StringVar(&path, "config_path", "./cmd/presence/config.yaml", "path to config file")
	flag.Parse()

	cfg := config.MustReadConfig[config.PresenceConfig](path)
	appContainer, err := presence.NewApp(cfg)
	if err != nil {
		log.Fatalf("can not create presence app: %v", err)
	}

	log.Fatal(grpc.Run(appContainer, cfg))
}
