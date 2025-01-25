package port

import (
	"context"

	"github.com/RezaMokaram/chapp/internal/chat/domain"
)

type Pubsub interface {
	SubscribeToMessages(context.Context, domain.RoomId, domain.UserId) (<-chan domain.Message, error)
	UnSubscribe(ctx context.Context) error
	PublishMessage(ctx context.Context, message domain.Message) error
}
