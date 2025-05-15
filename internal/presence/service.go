package presence

import (
	"context"
	"errors"

	"github.com/rezamokaram/dchat/internal/presence/domain"
	"github.com/rezamokaram/dchat/internal/presence/port"
)

var (
	ErrPresenceOnCreate           = errors.New("error on creating new presence")
	ErrPresenceCreationValidation = errors.New("validation failed")
	ErrPresenceNotFound           = errors.New("presence not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// TODO: lock & validation
func (ps *service) SetUserPresence(ctx context.Context, user domain.User) error {
	return ps.repo.SetUserPresence(ctx, user)
}
func (ps *service) DeleteUserPresence(ctx context.Context, userId domain.UserId) error {
	return ps.repo.DeleteUserPresence(ctx, userId)
}
func (ps *service) GetUsersByFilter(ctx context.Context, filter domain.UserFilter) ([]domain.User, error) {
	return ps.repo.GetUsersByFilter(ctx, filter)
}
func (ps *service) GetRoomByFilter(ctx context.Context, filter domain.RoomFilter) (*domain.Room, error) {
	return ps.repo.GetRoomByFilter(ctx, filter)
}
