package mapper

import (
	"github.com/RezaMokaram/chapp/internal/user/domain"
	"github.com/RezaMokaram/chapp/pkg/adapters/room_storage/types"
	"github.com/google/uuid"
)

func UserDomain2Storage(userDomain domain.User) (*types.User, error) {
	res := &types.User{
		CreatedAt: userDomain.CreatedAt,
		DeletedAt: ToDeletedAt(userDomain.DeletedAt),
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		Username:  userDomain.Username,
	}

	var err error
	res.ID, err = uuid.Parse(string(userDomain.ID))
	return res, err
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
