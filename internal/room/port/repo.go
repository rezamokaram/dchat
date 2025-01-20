package port

import (
	"context"

	"github.com/RezaMokaram/chapp/internal/room/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.Room) (domain.RoomID, error)
	GetByFilter(ctx context.Context, filter *domain.RoomFilter) (*domain.Room, error)
}
