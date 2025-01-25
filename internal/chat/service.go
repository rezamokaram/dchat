package chat

import (
	"context"
	"errors"

	"github.com/RezaMokaram/chapp/config"
	"github.com/RezaMokaram/chapp/internal/chat/domain"
	"github.com/RezaMokaram/chapp/internal/chat/port"
)

var (
	ErrNoNewMessages = errors.New("there is no new messages")
)

type service struct {
	cfg    config.ChatConfig
	pubsub port.Pubsub
}

func NewService(cfg config.ChatConfig, pubsub port.Pubsub) port.Service {
	return &service{
		cfg:    cfg,
		pubsub: pubsub,
	}
}

func (s *service) Send(ctx context.Context, message domain.Message) error {
	// todo update presence of user

	return s.pubsub.PublishMessage(ctx, message)
}

func (s *service) Receiver(ctx context.Context, roomId domain.RoomId, userId domain.UserId) (<-chan domain.Message, error) {
	return s.pubsub.SubscribeToMessages(ctx, roomId, userId)
}
