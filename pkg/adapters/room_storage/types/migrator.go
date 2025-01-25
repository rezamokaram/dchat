package types

import "gorm.io/gorm"

// I don't like code first solutions, but now for a force task i will use it.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Room{},
		&User{},
	)
}
