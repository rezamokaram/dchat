package port

import (
	"context"

	"github.com/rezamokaram/dchat/internal/presence/domain"
)

type Repo interface {
	SetUserPresence(ctx context.Context, user domain.User) error
	DeleteUserPresence(ctx context.Context, user domain.UserId) error
	GetUsersByFilter(ctx context.Context, filter domain.UserFilter) ([]domain.User, error)
	GetRoomByFilter(ctx context.Context, filter domain.RoomFilter) (*domain.Room, error)
}
