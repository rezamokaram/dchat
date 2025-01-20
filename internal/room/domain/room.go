package domain

type (
	RoomID string
)

type Room struct {
	ID      RoomID
	OwnerId string
	Name    string
}

type RoomFilter struct {
	ID RoomID
}
