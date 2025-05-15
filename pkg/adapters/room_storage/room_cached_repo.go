package room_storage

import (
	"context"
	"log"

	"github.com/rezamokaram/dchat/internal/room/domain"
	roomPort "github.com/rezamokaram/dchat/internal/room/port"
	"github.com/rezamokaram/dchat/pkg/cache"
)

type roomCachedRepo struct {
	repo     roomPort.Repo
	provider cache.Provider
}

func (r *roomCachedRepo) Create(ctx context.Context, roomDomain domain.Room) (domain.RoomID, error) {
	uId, err := r.repo.Create(ctx, roomDomain)
	if err != nil {
		return "", err
	}
	roomDomain.ID = uId

	oc := cache.NewJsonObjectCacher[*domain.Room](r.provider)
	if err := oc.Set(ctx, r.roomFilterKey(&domain.RoomFilter{
		ID: uId,
	}), 0, &roomDomain); err != nil {
		log.Println("error on caching (SET) room with id :", uId)
	}

	return uId, nil
}

func (r *roomCachedRepo) roomFilterKey(filter *domain.RoomFilter) string {
	return "rooms." + string(filter.ID)
}

func (r *roomCachedRepo) GetByFilter(ctx context.Context, filter *domain.RoomFilter) (*domain.Room, error) {
	oc := cache.NewJsonObjectCacher[*domain.Room](r.provider)

	key := r.roomFilterKey(filter)
	dRoom, err := oc.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if dRoom != nil && len(dRoom.ID) > 0 {
		return dRoom, nil
	}

	dRoom, err = r.repo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	if dRoom == nil {
		return nil, nil
	}

	if err := oc.Set(ctx, key, 0, dRoom); err != nil {
		log.Printf("error on caching (SET) room with filter : %+v", *filter)
	}

	return dRoom, nil
}
