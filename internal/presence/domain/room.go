package domain

import "github.com/google/uuid"

type (
	RoomId uuid.UUID
)

type Room struct {
	ID RoomId
	Users map[UserId]User
}

type RoomFilter struct {
	ID RoomId
}