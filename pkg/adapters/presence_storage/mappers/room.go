package mappers

import (
	"github.com/google/uuid"
	"github.com/rezamokaram/dchat/internal/presence/domain"
	"github.com/rezamokaram/dchat/pkg/adapters/presence_storage/types"
)

func RoomDomain2Storage(roomDomain domain.Room) *types.Room {
	room := &types.Room{
		ID: uuid.UUID(roomDomain.ID),
	}
	for _, usr := range roomDomain.Users {
		val := *UserDomain2Storage(usr)
		room.Users[val.ID] = val
	}
	return room
}

func RoomStorage2Domain(room types.Room) *domain.Room {
	roomDomain := &domain.Room{
		ID:    domain.RoomId(room.ID),
		Users: make(map[domain.UserId]domain.User),
	}

	if room.Users == nil {
		return roomDomain
	}

	for _, usr := range room.Users {
		val := *UserStorage2Domain(usr)
		roomDomain.Users[val.ID] = val
	}
	return roomDomain
}
