package mapper

import (
	"github.com/RezaMokaram/chapp/internal/room/domain"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/types"
	"github.com/google/uuid"
)

func RoomDomain2Storage(roomDomain domain.Room) *types.Room {
	res := &types.Room{
		Name: roomDomain.Name,
	}
	res.ID, _ = uuid.Parse(string(roomDomain.ID))
	res.OwnerId, _ = uuid.Parse(string(roomDomain.OwnerId))
	return res
}

func RoomStorage2Domain(room types.Room) *domain.Room {
	return &domain.Room{
		ID:      domain.RoomID(room.ID.String()),
		OwnerId: room.OwnerId.String(),
		Name:    room.Name,
	}
}
