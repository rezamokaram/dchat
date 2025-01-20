package room

import (
	"context"
	"fmt"

	"github.com/RezaMokaram/chapp/config"
	"github.com/RezaMokaram/chapp/internal/room"
	roomPort "github.com/RezaMokaram/chapp/internal/room/port"
	"github.com/RezaMokaram/chapp/internal/user"
	userPort "github.com/RezaMokaram/chapp/internal/user/port"
	storage "github.com/RezaMokaram/chapp/pkg/adapters/room_storage"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/types"
	"github.com/RezaMokaram/chapp/pkg/cache"
	"github.com/RezaMokaram/chapp/pkg/postgres"

	// "github.com/go-co-op/gocron/v2"

	redisAdapter "github.com/RezaMokaram/chapp/pkg/adapters/cache"

	"gorm.io/gorm"

	appCtx "github.com/RezaMokaram/chapp/pkg/context"
)

type app struct {
	db            *gorm.DB
	cfg           config.RoomConfig
	userService   userPort.Service
	roomService   roomPort.Service
	redisProvider cache.Provider
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db, true, a.redisProvider))
}

func (a *app) RoomService(ctx context.Context) roomPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.roomService == nil {
			a.roomService = a.roomServiceWithDB(a.db)
		}
		return a.roomService
	}

	return a.roomServiceWithDB(db)
}

func (a *app) roomServiceWithDB(db *gorm.DB) roomPort.Service {
	return room.NewService(storage.NewRoomRepo(db, true, a.redisProvider))
}

func (a *app) Config() config.RoomConfig {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.Postgres.User,
		Pass:   a.cfg.Postgres.Password,
		Host:   a.cfg.Postgres.Host,
		Port:   a.cfg.Postgres.Port,
		DBName: a.cfg.Postgres.DB,
		Schema: a.cfg.Postgres.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setRedis() {
	a.redisProvider = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func NewApp(cfg config.RoomConfig) (RoomApp, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	if err := types.Migrate(a.db); err != nil {
		return nil, err
	}

	a.setRedis()

	return a, nil //a.registerOutboxHandlers()
}

func NewMustApp(cfg config.RoomConfig) RoomApp {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
