package port

import (
	"context"

	"github.com/RezaMokaram/chapp/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	GetUserByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error)
}
