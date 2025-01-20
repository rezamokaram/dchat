package grpc

import (
	"context"
	// "log"

	"github.com/RezaMokaram/chapp/api/handler/common"
	"github.com/RezaMokaram/chapp/api/service"
	app "github.com/RezaMokaram/chapp/app/room"
	"github.com/RezaMokaram/chapp/config"
)

func roomServiceGetter(appContainer app.RoomApp, cfg config.RoomConfig) common.ServiceGetter[*service.RoomService] {
	return func(ctx context.Context) *service.RoomService {
		res := service.NewRoomService(appContainer.UserService(ctx), appContainer.RoomService(ctx))
		if res == nil {
			panic("service getter problem!")
		}
		return res
	}
}
