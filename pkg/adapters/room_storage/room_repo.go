package room_storage

import (
	"context"
	"errors"

	"github.com/rezamokaram/dchat/internal/room/domain"
	"github.com/rezamokaram/dchat/internal/room/port"
	"github.com/rezamokaram/dchat/pkg/adapters/room_storage/mapper"
	"github.com/rezamokaram/dchat/pkg/adapters/room_storage/types"
	"github.com/rezamokaram/dchat/pkg/cache"
	"gorm.io/gorm"
)

type roomRepo struct {
	db *gorm.DB
}

func NewRoomRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	repo := &roomRepo{db}
	if !cached {
		return repo
	}

	return &roomCachedRepo{
		repo:     repo,
		provider: provider,
	}
}

func (r *roomRepo) Create(ctx context.Context, roomDomain domain.Room) (domain.RoomID, error) {
	room := mapper.RoomDomain2Storage(roomDomain)
	err := r.db.Table("rooms").WithContext(ctx).Create(room).Error
	return domain.RoomID(room.ID.String()), err
}

func (r *roomRepo) GetByFilter(ctx context.Context, filter *domain.RoomFilter) (*domain.Room, error) {
	var room types.Room

	q := r.db.Table("room").Debug().WithContext(ctx)

	if len(filter.ID) > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	err := q.First(&room).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if len(room.ID) == 0 {
		return nil, nil
	}

	return mapper.RoomStorage2Domain(room), nil
}
