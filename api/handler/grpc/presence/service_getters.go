package presence

import (
	"context"

	"github.com/rezamokaram/dchat/api/handler/common"
	"github.com/rezamokaram/dchat/api/service"
	app "github.com/rezamokaram/dchat/app/presence"
)

func presenceServiceGetter(appContainer app.PresenceApp) common.ServiceGetter[*service.PresenceService] {
	return func(ctx context.Context) *service.PresenceService {
		res := service.NewPresenceService(appContainer.PresenceService(ctx))
		if res == nil {
			panic("service getter problem!")
		}
		return res
	}
}
