package mapper

import (
	"github.com/google/uuid"
	"github.com/rezamokaram/dchat/internal/user/domain"
	"github.com/rezamokaram/dchat/pkg/adapters/room_storage/types"
)

func UserDomain2Storage(userDomain domain.User) *types.User {
	res := &types.User{
		CreatedAt: userDomain.CreatedAt,
		DeletedAt: ToDeletedAt(userDomain.DeletedAt),
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		Username:  userDomain.Username,
	}

	res.ID, _ = uuid.Parse(string(userDomain.ID))
	return res
}

func UserStorage2Domain(user types.User) *domain.User {
	return &domain.User{
		ID:        domain.UserID(user.ID.String()),
		CreatedAt: user.CreatedAt,
		DeletedAt: user.DeletedAt.Time,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}
}
