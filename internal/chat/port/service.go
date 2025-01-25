package port

import (
	"context"

	"github.com/RezaMokaram/chapp/internal/chat/domain"
)

type Service interface {
	Send(ctx context.Context, message domain.Message) error
	Receiver(ctx context.Context, roomId domain.RoomId, userId domain.UserId) (<-chan domain.Message, error)
}
