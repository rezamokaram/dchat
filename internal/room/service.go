package room

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rezamokaram/dchat/internal/room/domain"
	"github.com/rezamokaram/dchat/internal/room/port"
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

func (s *service) CreateRoom(ctx context.Context, room domain.Room) (domain.RoomID, error) {
	if err := room.Validate(); err != nil {
		return "", fmt.Errorf("%w %w", ErrRoomCreationValidation, err)
	}

	roomID, err := s.repo.Create(ctx, room)
	if err != nil {
		log.Println("error on creating new room : ", err.Error())
		return "", ErrRoomOnCreate
	}

	return roomID, nil
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
