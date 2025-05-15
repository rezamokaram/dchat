package main

import (
	"flag"
	"fmt"
	"log"

	grpc "github.com/rezamokaram/dchat/api/handler/grpc/room"
	"github.com/rezamokaram/dchat/app/room"
	"github.com/rezamokaram/dchat/config"
)

func main() {
	var path string
	flag.StringVar(&path, "config_path", "./cmd/room/config.yaml", "path to config file")
	flag.Parse()

	cfg := config.MustReadConfig[config.RoomConfig](path)

	fmt.Println(cfg)
	appContainer := room.NewMustApp(cfg)

	log.Fatal(grpc.Run(appContainer, cfg))
}
