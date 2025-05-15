package chat

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/rezamokaram/dchat/config"
	chatPort "github.com/rezamokaram/dchat/internal/chat/port"
)

type ChatApp interface {
	ChatService(ctx context.Context) chatPort.Service
	Nats() *nats.Conn
	Config() config.ChatConfig
}
