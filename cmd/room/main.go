package main

import (
	"flag"
	"fmt"

	"github.com/RezaMokaram/chapp/config"
)

func main() {
	var path string
	flag.StringVar(&path, "config_path", "./cmd/room/config.yaml", "path to config file")
	flag.Parse()

	cfg := config.MustReadConfig[config.RoomConfig](path)

	fmt.Println(cfg)
}
