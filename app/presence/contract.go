package presence

import (
	"context"

	"github.com/RezaMokaram/chapp/config"
	presencePort "github.com/RezaMokaram/chapp/internal/presence/port"
	client "go.etcd.io/etcd/client/v3"
)

type PresenceApp interface {
	PresenceService(ctx context.Context) presencePort.Service
	DB() *client.Client
	Config() config.PresenceConfig
}
