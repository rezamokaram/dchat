package in_memory_kv_storage

import (
	"encoding/json"

	"github.com/google/uuid"
)

func getUserKey(userId uuid.UUID) string {
	return "user." + userId.String()
}

func getRoomKey(roomId uuid.UUID) string {
	return "room." + roomId.String()
}

func struct2String[T any](in T) (string, error) {
	jsonData, err := json.Marshal(in)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
