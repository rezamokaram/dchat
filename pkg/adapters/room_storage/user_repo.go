package room_storage

import (
	"context"
	"errors"

	"github.com/RezaMokaram/chapp/internal/user/domain"
	"github.com/RezaMokaram/chapp/internal/user/port"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/mapper"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/types"
	"github.com/RezaMokaram/chapp/pkg/cache"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	repo := &userRepo{db}
	if !cached {
		return repo
	}

	return &userCachedRepo{
		repo:     repo,
		provider: provider,
	}
}

func (r *userRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	user, err := mapper.UserDomain2Storage(userDomain)
	if err != nil {
		return "", err
	}

	return domain.UserID(user.ID.String()), r.db.Table("users").WithContext(ctx).Create(user).Error
}

func (r *userRepo) GetByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	var user types.User

	q := r.db.Table("users").Debug().WithContext(ctx)

	if len(filter.ID) > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if len(filter.Username) > 0 {
		q = q.Where("username = ?", filter.Username)
	}

	err := q.First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if len(user.ID) == 0 {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}
