package mapper

import (
	"github.com/google/uuid"
	"github.com/rezamokaram/dchat/internal/room/domain"
	"github.com/rezamokaram/dchat/pkg/adapters/room_storage/types"
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
