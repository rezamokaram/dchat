package room_storage

import (
	"context"
	"errors"

	"github.com/RezaMokaram/chapp/internal/room/domain"
	"github.com/RezaMokaram/chapp/internal/room/port"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/mapper"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/types"
	"github.com/RezaMokaram/chapp/pkg/cache"
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

func (r *roomRepo) Create(ctx context.Context, userDomain domain.Room) (domain.RoomID, error) {
	user, err := mapper.RoomDomain2Storage(userDomain)
	if err != nil {
		return "", err
	}

	return domain.RoomID(user.ID.String()), r.db.Table("users").WithContext(ctx).Create(user).Error
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
