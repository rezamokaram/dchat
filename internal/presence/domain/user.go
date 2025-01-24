package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserId uuid.UUID
)
type User struct {
	ID        UserId
	RoomID    RoomId
	Status    uint
	UpdatedAt time.Time
}

type UserFilter struct {
	ID        UserId
	RoomID    RoomId
	Status    bool
	UpdatedAt time.Time
}

func (r UserId) ToString() string {
	return uuid.UUID(r).String()
}
