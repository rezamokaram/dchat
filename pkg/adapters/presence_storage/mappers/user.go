package mappers

import (
	"github.com/RezaMokaram/chapp/internal/presence/domain"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_storage/types"
	"github.com/google/uuid"
)

func UserDomain2Storage(userDomain domain.User) *types.User {
	return &types.User{
		ID:       uuid.UUID(userDomain.ID),
		RoomID:    uuid.UUID(userDomain.RoomID),
		Status:    userDomain.Status,
		UpdatedAt: userDomain.UpdatedAt,
	}
}

func UserStorage2Domain(user types.User) *domain.User {
	return &domain.User{
		ID:       domain.UserId(user.ID),
		RoomID:    domain.RoomId(user.RoomID),
		Status:    user.Status,
		UpdatedAt: user.UpdatedAt,
	}
}