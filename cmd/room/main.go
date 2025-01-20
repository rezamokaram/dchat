package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/RezaMokaram/chapp/api/handler/grpc"
	"github.com/RezaMokaram/chapp/app/room"
	"github.com/RezaMokaram/chapp/config"
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
