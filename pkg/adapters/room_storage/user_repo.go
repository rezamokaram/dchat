package room_storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/rezamokaram/dchat/internal/user/domain"
	"github.com/rezamokaram/dchat/internal/user/port"
	"github.com/rezamokaram/dchat/pkg/adapters/room_storage/mapper"
	"github.com/rezamokaram/dchat/pkg/adapters/room_storage/types"
	"github.com/rezamokaram/dchat/pkg/cache"

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
	fmt.Println("REZA: ", userDomain)
	user := mapper.UserDomain2Storage(userDomain)
	err := r.db.Table("users").WithContext(ctx).Create(user).Error
	return domain.UserID(user.ID.String()), err
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
