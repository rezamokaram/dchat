package grpc

import (
	"context"

	"github.com/RezaMokaram/chapp/api/handler/common"
	"github.com/RezaMokaram/chapp/api/service"
	app "github.com/RezaMokaram/chapp/app/room"
	"github.com/RezaMokaram/chapp/config"
)

// user service transient instance handler
func userServiceGetter(appContainer app.RoomApp, cfg config.RoomConfig) common.ServiceGetter[*service.RoomService] {
	return func(ctx context.Context) *service.RoomService {
		return service.NewRoomService(appContainer.UserService(ctx), appContainer.RoomService(ctx))
	}
}
