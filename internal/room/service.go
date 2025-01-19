package room

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/RezaMokaram/chapp/internal/room/domain"
	"github.com/RezaMokaram/chapp/internal/room/port"
)

var (
	ErrRoomOnCreate           = errors.New("error on creating new room")
	ErrRoomCreationValidation = errors.New("validation failed")
	ErrRoomNotFound           = errors.New("room not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateRoom(ctx context.Context, user domain.Room) (domain.RoomID, error) {
	if err := user.Validate(); err != nil {
		return "", fmt.Errorf("%w %w", ErrRoomCreationValidation, err)
	}

	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		log.Println("error on creating new user : ", err.Error())
		return "", ErrRoomOnCreate
	}

	return userID, nil
}

func (s *service) GetRoomByFilter(ctx context.Context, filter *domain.RoomFilter) (*domain.Room, error) {
	room, err := s.repo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	if room == nil {
		return nil, ErrRoomNotFound
	}

	return room, nil
}
