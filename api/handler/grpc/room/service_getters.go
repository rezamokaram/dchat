package room

import (
	"context"
	// "log"

	"github.com/rezamokaram/dchat/api/handler/common"
	"github.com/rezamokaram/dchat/api/service"
	app "github.com/rezamokaram/dchat/app/room"
)

func roomServiceGetter(appContainer app.RoomApp) common.ServiceGetter[*service.RoomService] {
	return func(ctx context.Context) *service.RoomService {
		res := service.NewRoomService(appContainer.UserService(ctx), appContainer.RoomService(ctx))
		if res == nil {
			panic("service getter problem!")
		}
		return res
	}
}
