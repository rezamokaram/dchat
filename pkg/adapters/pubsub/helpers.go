package pubsub

import (
	"fmt"
)

func getSubjectName(roomId string) string {
	return fmt.Sprintf("room.%s", roomId)
}
