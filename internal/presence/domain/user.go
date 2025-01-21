package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserId uuid.UUID
)
type User struct {
	ID UserId
	RoomID RoomId
	Status uint
	UpdatedAt time.Time
}

type UserFilter struct {
	ID uuid.UUID
	RoomID uuid.UUID
	Status bool
	UpdatedAt time.Time
}