package chat

import (
	"context"

	"github.com/RezaMokaram/chapp/config"
	chatPort "github.com/RezaMokaram/chapp/internal/chat/port"
	"github.com/nats-io/nats.go"
)

type ChatApp interface {
	ChatService(ctx context.Context) chatPort.Service
	Nats() *nats.Conn
	Config() config.ChatConfig
}
