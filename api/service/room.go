package service

import (
	"context"

	"github.com/rezamokaram/dchat/api/pb"
	roomDomain "github.com/rezamokaram/dchat/internal/room/domain"
	roomPort "github.com/rezamokaram/dchat/internal/room/port"
	"github.com/rezamokaram/dchat/internal/user"
	userDomain "github.com/rezamokaram/dchat/internal/user/domain"
	userPort "github.com/rezamokaram/dchat/internal/user/port"
)

type RoomService struct {
	userSvc userPort.Service
	roomSvc roomPort.Service
}

func NewRoomService(
	userSvc userPort.Service,
	roomSvc roomPort.Service,
) *RoomService {
	return &RoomService{
		userSvc: userSvc,
		roomSvc: roomSvc,
	}
}

var (
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserOnCreate           = user.ErrUserOnCreate
	ErrUserNotFound           = user.ErrUserNotFound
)

func (s *RoomService) SignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	userID, err := s.userSvc.CreateUser(ctx, userDomain.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Username:  req.GetUsername(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UserSignUpResponse{
		Success: true,
		UserId:  string(userID),
	}, nil
}

func (s *RoomService) SignIn(ctx context.Context, req *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	user, err := s.userSvc.GetUserByFilter(ctx, &userDomain.UserFilter{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return &pb.UserSignInResponse{
		Success: true,
		UserId:  string(user.ID),
	}, nil
}

func (s *RoomService) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	roomId, err := s.roomSvc.CreateRoom(ctx, roomDomain.Room{
		OwnerId: req.UserId,
		Name:    req.RoomName,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomResponse{
		Success: true,
		RoomId:  string(roomId),
	}, nil
}
