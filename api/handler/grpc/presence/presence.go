package presence

import (
	"context"
	"errors"

	"github.com/RezaMokaram/chapp/api/handler/common"
	"github.com/RezaMokaram/chapp/api/handler/grpc/interceptors"
	"github.com/RezaMokaram/chapp/api/pb"
	"github.com/RezaMokaram/chapp/api/service"
	"github.com/RezaMokaram/chapp/app/presence"
)

type presenceServer struct {
	pb.UnimplementedPresenceServiceServer
	appContainer presence.PresenceApp
}

func (pr *presenceServer) SetUserPresence(ctx context.Context, in *pb.SetUserPresenceRequest) (*pb.SetUserPresenceResponse, error) {
	presenceSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.PresenceService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := presenceSvcGetter(ctx)
	return svc.SetUserPresenceData(ctx, in)
}

func (pr *presenceServer) DeleteUserPresence(ctx context.Context, in *pb.DeleteUserPresenceRequest) (*pb.DeleteUserPresenceResponse, error) {
	presenceSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.PresenceService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := presenceSvcGetter(ctx)
	return svc.DeleteUserPresenceData(ctx, in)
}

func (pr *presenceServer) GetRoomPresenceData(ctx context.Context, in *pb.GetRoomPresenceDataRequest) (*pb.GetRoomPresenceDataResponse, error) {
	presenceSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.PresenceService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := presenceSvcGetter(ctx)
	return svc.GetRoomPresenceData(ctx, in)
}

func (pr *presenceServer) GetUserPresence(ctx context.Context, in *pb.GetUserPresenceRequest) (*pb.GetUserPresenceResponse, error) {
	presenceSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.PresenceService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := presenceSvcGetter(ctx)
	return svc.GetUserPresence(ctx, in)
}

func (pr *presenceServer) GetUsersPresence(ctx context.Context, in *pb.GetUsersPresenceRequest) (*pb.GetUsersPresenceResponse, error) {
	presenceSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.PresenceService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := presenceSvcGetter(ctx)
	return svc.GetUsersPresence(ctx, in)
}
