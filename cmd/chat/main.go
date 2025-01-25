package main

import (
	"flag"
	"log"

	grpc "github.com/RezaMokaram/chapp/api/handler/grpc/chat"
	"github.com/RezaMokaram/chapp/app/chat"
	"github.com/RezaMokaram/chapp/config"
)

func main() {
	var path string
	flag.StringVar(&path, "config_path", "./cmd/chat/config.yaml", "path to config file")
	flag.Parse()
	cfg := config.MustReadConfig[config.ChatConfig](path)
	appContainer, err := chat.NewApp(cfg)
	if err != nil {
		log.Fatalf("can not create chat app: %v", err)
	}

	log.Fatal(grpc.Run(appContainer, cfg))
}
