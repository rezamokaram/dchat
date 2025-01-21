package types

import "github.com/google/uuid"


type Room struct {
	ID    uuid.UUID          `json:"id,omitempty"`
	Users map[uuid.UUID]User `json:"users,omitempty"`
}
