package room

import (
	"context"

	"github.com/RezaMokaram/chapp/config"
	roomPort "github.com/RezaMokaram/chapp/internal/room/port"
	userPort "github.com/RezaMokaram/chapp/internal/user/port"

	"gorm.io/gorm"
)

type RoomApp interface {
	UserService(ctx context.Context) userPort.Service
	RoomService(ctx context.Context) roomPort.Service
	DB() *gorm.DB
	Config() config.RoomConfig
}
