package presence

import (
	"context"

	"github.com/RezaMokaram/chapp/config"
	"github.com/RezaMokaram/chapp/internal/presence"
	presencePort "github.com/RezaMokaram/chapp/internal/presence/port"
	storage "github.com/RezaMokaram/chapp/pkg/adapters/presence_storage"
	"github.com/RezaMokaram/chapp/pkg/etcd"
	client "go.etcd.io/etcd/client/v3"
)

type app struct {
	cfg config.PresenceConfig
	db  *client.Client
}

func NewApp(cfg config.PresenceConfig) (PresenceApp, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil //a.registerOutboxHandlers()
}

func (a *app) setDB() error {
	db, err := etcd.NewEtcdClient(a.cfg.Etcd.Hosts)

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) DB() *client.Client {
	return a.db
}

func (a *app) PresenceService(ctx context.Context) presencePort.Service {
	return presence.NewService(storage.NewPresenceRepo(a.db, a.cfg.Etcd.TTL))
}

func (a *app) Config() config.PresenceConfig {
	return a.cfg
}
