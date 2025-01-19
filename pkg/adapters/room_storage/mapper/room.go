package mapper

import (
	"github.com/RezaMokaram/chapp/internal/room/domain"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/types"
	"github.com/google/uuid"
)

func RoomDomain2Storage(userDomain domain.Room) (*types.Room, error) {
	res := &types.Room{}
	var err error
	res.OwnerId, err = uuid.Parse(string(userDomain.OwnerId))
	if err != nil {
		return nil, err
	}

	res.ID, err = uuid.Parse(string(userDomain.ID))
	return res, err
}

func RoomStorage2Domain(user types.Room) *domain.Room {
	return &domain.Room{
		ID:      domain.RoomID(user.ID.String()),
		OwnerId: user.OwnerId.String(),
	}
}
