package presence

import (
	"context"

	"github.com/RezaMokaram/chapp/api/handler/common"
	"github.com/RezaMokaram/chapp/api/service"
	app "github.com/RezaMokaram/chapp/app/presence"
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
