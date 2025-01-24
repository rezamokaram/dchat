package types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	RoomID    uuid.UUID `json:"room_id,omitempty"`
	Status    uint      `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
