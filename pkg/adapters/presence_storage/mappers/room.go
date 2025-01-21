package mappers

import (
	"github.com/RezaMokaram/chapp/internal/presence/domain"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_storage/types"
	"github.com/google/uuid"
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
		ID: domain.RoomId(room.ID),
	}
	for _, usr := range room.Users {
		val := *UserStorage2Domain(usr)
		roomDomain.Users[val.ID] = val
	}
	return roomDomain
}