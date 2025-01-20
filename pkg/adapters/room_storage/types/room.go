package types

import (
	// "time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type Room struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerId uuid.UUID `json:"owner_id"`
	Name    string    `json:"name"`
}
