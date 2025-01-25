package port

import (
	"context"
)

type PresenceClient interface {
	SetUserPresence(ctx context.Context, roomId string, userId string) error
}
