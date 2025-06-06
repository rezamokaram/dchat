package room

import (
	"context"
	"errors"

	"github.com/rezamokaram/dchat/api/handler/common"
	"github.com/rezamokaram/dchat/api/handler/grpc/interceptors"
	"github.com/rezamokaram/dchat/api/pb"
	"github.com/rezamokaram/dchat/api/service"
	"github.com/rezamokaram/dchat/app/room"
)

type roomServer struct {
	pb.UnimplementedRoomServiceServer
	appContainer room.RoomApp
}

func (s *roomServer) SignInUser(ctx context.Context, in *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	roomSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.RoomService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := roomSvcGetter(ctx)
	return svc.SignIn(ctx, in)
}

func (s *roomServer) SignUpUser(ctx context.Context, in *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	roomSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.RoomService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := roomSvcGetter(ctx)
	return svc.SignUp(ctx, in)
}

func (s *roomServer) CreateRoom(ctx context.Context, in *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	roomSvcGetter, exist := ctx.Value(interceptors.SvcContextKey).(common.ServiceGetter[*service.RoomService])
	if !exist {
		return nil, errors.New("service is not provided")
	}

	svc := roomSvcGetter(ctx)
	return svc.CreateRoom(ctx, in)
}
