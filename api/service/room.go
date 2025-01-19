package service

import (
	"context"

	"github.com/RezaMokaram/chapp/api/pb"
	roomPort "github.com/RezaMokaram/chapp/internal/room/port"
	"github.com/RezaMokaram/chapp/internal/user"
	"github.com/RezaMokaram/chapp/internal/user/domain"
	userPort "github.com/RezaMokaram/chapp/internal/user/port"
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
	userID, err := s.userSvc.CreateUser(ctx, domain.User{
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
	user, err := s.userSvc.GetUserByFilter(ctx, &domain.UserFilter{
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
