package presence

import (
	"context"

	"github.com/rezamokaram/dchat/config"
	"github.com/rezamokaram/dchat/internal/presence"
	presencePort "github.com/rezamokaram/dchat/internal/presence/port"
	storage "github.com/rezamokaram/dchat/pkg/adapters/presence_storage"
	"github.com/rezamokaram/dchat/pkg/etcd"
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
