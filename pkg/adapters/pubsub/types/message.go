package types

type Message struct {
	UserId  string `json:"user_id,omitempty"`
	RoomId  string `json:"room_id,omitempty"`
	Content string `json:"content,omitempty"`
}
