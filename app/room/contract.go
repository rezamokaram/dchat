package room

import (
	"context"

	"github.com/rezamokaram/dchat/config"
	roomPort "github.com/rezamokaram/dchat/internal/room/port"
	userPort "github.com/rezamokaram/dchat/internal/user/port"

	"gorm.io/gorm"
)

type RoomApp interface {
	UserService(ctx context.Context) userPort.Service
	RoomService(ctx context.Context) roomPort.Service
	DB() *gorm.DB
	Config() config.RoomConfig
}
