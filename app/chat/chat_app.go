package chat

import (
	"context"

	"github.com/RezaMokaram/chapp/config"
	"github.com/RezaMokaram/chapp/internal/chat"
	chatPort "github.com/RezaMokaram/chapp/internal/chat/port"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_client"
	"github.com/RezaMokaram/chapp/pkg/adapters/pubsub"
	pkgNats "github.com/RezaMokaram/chapp/pkg/nats"
	"github.com/nats-io/nats.go"
)

type app struct {
	cfg config.ChatConfig
	nc  *nats.Conn
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

	app.prClient, err = presence_client.NewPresenceClient(cfg.Chat.Phost +":" + cfg.Chat.Pport)
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
