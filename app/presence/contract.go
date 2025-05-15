package presence

import (
	"context"

	"github.com/rezamokaram/dchat/config"
	presencePort "github.com/rezamokaram/dchat/internal/presence/port"
	client "go.etcd.io/etcd/client/v3"
)

type PresenceApp interface {
	PresenceService(ctx context.Context) presencePort.Service
	DB() *client.Client
	Config() config.PresenceConfig
}
