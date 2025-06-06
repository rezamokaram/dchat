package room_storage

import (
	"context"
	"log"

	"github.com/rezamokaram/dchat/internal/user/domain"
	userPort "github.com/rezamokaram/dchat/internal/user/port"
	"github.com/rezamokaram/dchat/pkg/cache"
)

type userCachedRepo struct {
	repo     userPort.Repo
	provider cache.Provider
}

func (r *userCachedRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	uId, err := r.repo.Create(ctx, userDomain)
	if err != nil {
		return "", err
	}
	userDomain.ID = uId

	oc := cache.NewJsonObjectCacher[*domain.User](r.provider)
	if err := oc.Set(ctx, r.userFilterKey(&domain.UserFilter{
		ID: uId,
	}), 0, &userDomain); err != nil {
		log.Println("error on caching (SET) user with id :", uId)
	}

	return uId, nil
}

func (r *userCachedRepo) userFilterKey(filter *domain.UserFilter) string {
	return "users." + string(filter.ID) + "." + filter.Username
}

func (r *userCachedRepo) GetByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	oc := cache.NewJsonObjectCacher[*domain.User](r.provider)

	key := r.userFilterKey(filter)
	dUser, err := oc.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if dUser != nil && len(dUser.ID) > 0 {
		return dUser, nil
	}

	dUser, err = r.repo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	if dUser == nil {
		return nil, nil
	}

	if err := oc.Set(ctx, key, 0, dUser); err != nil {
		log.Printf("error on caching (SET) user with filter : %+v", *filter)
	}

	return dUser, nil
}
