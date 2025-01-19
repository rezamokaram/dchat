package domain

type (
	RoomID string
)

type Room struct {
	ID      RoomID
	OwnerId string
}

type RoomFilter struct {
	ID RoomID
}
