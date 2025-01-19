package types

import (
	// "time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type Room struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OwnerId uuid.UUID
}
