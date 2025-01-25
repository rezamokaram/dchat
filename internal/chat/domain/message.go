package domain

import "github.com/google/uuid"

type (
	UserId uuid.UUID
	RoomId uuid.UUID
)

type Message struct {
	UserId  UserId
	RoomId  RoomId
	Content string
	Filled  bool
}

func (r RoomId) ToString() string {
	return uuid.UUID(r).String()
}

func (u UserId) ToString() string {
	return uuid.UUID(u).String()
}
