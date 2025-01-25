package in_memory_kv_storage

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func getUserKeyPrefix() string {
	return "user."
}

func getUserKey(userId uuid.UUID) string {
	if userId == uuid.Nil {
		return getUserKeyPrefix()
	}
	return fmt.Sprintf("%s%s", getUserKeyPrefix(), userId.String())
}

func getRoomKeyPrefix() string {
	return "room."
}

func getRoomKey(roomId uuid.UUID) string {
	if roomId == uuid.Nil {
		return getRoomKeyPrefix()
	}
	return fmt.Sprintf("%s%s", getRoomKeyPrefix(), roomId.String())
}

func struct2String[T any](in T) (string, error) {
	jsonData, err := json.Marshal(in)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
