package port

import (
	"context"

	"github.com/rezamokaram/dchat/internal/room/domain"
)

type Service interface {
	CreateRoom(ctx context.Context, user domain.Room) (domain.RoomID, error)
	GetRoomByFilter(ctx context.Context, filter *domain.RoomFilter) (*domain.Room, error)
}
