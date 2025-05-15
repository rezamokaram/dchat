package chat

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/rezamokaram/dchat/config"
	"github.com/rezamokaram/dchat/internal/chat"
	chatPort "github.com/rezamokaram/dchat/internal/chat/port"
	"github.com/rezamokaram/dchat/pkg/adapters/presence_client"
	"github.com/rezamokaram/dchat/pkg/adapters/pubsub"
	pkgNats "github.com/rezamokaram/dchat/pkg/nats"
)

type app struct {
	cfg      config.ChatConfig
	nc       *nats.Conn
	prClient chatPort.PresenceClient
}

func NewApp(cfg config.ChatConfig) (ChatApp, error) {
	app := &app{
		cfg: cfg,
	}

	var err error
	app.nc, err = pkgNats.NewNatsClient(cfg.Nats.Host)
	if err != nil {
		return nil, err
	}

	app.prClient, err = presence_client.NewPresenceClient(cfg.Chat.Phost + ":" + cfg.Chat.Pport)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *app) ChatService(ctx context.Context) chatPort.Service {
	return chat.NewService(a.cfg, pubsub.NewPubSub(a.nc), a.prClient)
}

func (a *app) Nats() *nats.Conn {
	return a.nc
}

func (a *app) Config() config.ChatConfig {
	return a.cfg
}
