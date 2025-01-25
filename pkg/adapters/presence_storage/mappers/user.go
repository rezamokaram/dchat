package mappers

import (
	"log"

	"github.com/RezaMokaram/chapp/internal/presence/domain"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_storage/types"
	"github.com/google/uuid"
)

func UserDomain2Storage(userDomain domain.User) *types.User {
	log.Printf("\n\n\n MAPPERS storageUser: %v, domainUser: %v\n\n\n", uuid.UUID(userDomain.RoomID).String(), userDomain.RoomID.ToString())
	return &types.User{
		ID:        uuid.UUID(userDomain.ID),
		RoomID:    uuid.UUID(userDomain.RoomID),
		Status:    userDomain.Status,
		UpdatedAt: userDomain.UpdatedAt,
	}
}

func UserStorage2Domain(user types.User) *domain.User {
	res := &domain.User{
		ID:        domain.UserId(user.ID),
		RoomID:    domain.RoomId(user.RoomID),
		Status:    user.Status,
		UpdatedAt: user.UpdatedAt,
	}
	log.Printf("\n\n\n MPPPPEEERRSS storageUser: %v %v, domainUser: %v\n\n\n", user.RoomID, user.RoomID, res.RoomID.ToString())
	return res
}
