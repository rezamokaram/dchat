package domain

import (
	"time"
)

type (
	UserID string
)

type User struct {
	ID        UserID
	CreatedAt time.Time
	DeletedAt time.Time
	FirstName string
	LastName  string
	Username  string
}

type UserFilter struct {
	ID       UserID
	Username string
}
