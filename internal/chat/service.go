package chat

import (
	"context"
	"errors"
	"log"

	"github.com/RezaMokaram/chapp/config"
	"github.com/RezaMokaram/chapp/internal/chat/domain"
	"github.com/RezaMokaram/chapp/internal/chat/port"
)

var (
	ErrNoNewMessages = errors.New("there is no new messages")
)

type service struct {
	cfg      config.ChatConfig
	pubsub   port.Pubsub
	prClient port.PresenceClient
}

func NewService(cfg config.ChatConfig, pubsub port.Pubsub, prClient port.PresenceClient) port.Service {
	return &service{
		cfg:      cfg,
		pubsub:   pubsub,
		prClient: prClient,
	}
}

func (s *service) Send(ctx context.Context, message domain.Message) error {
	err := s.pubsub.PublishMessage(ctx, message)

	if err == nil {
		err = s.prClient.SetUserPresence(ctx, message.RoomId.ToString(), message.UserId.ToString())
		if err != nil {
			log.Printf("can not update user presence: %v\n", err)
		}
	} else {
		return err
	}

	return nil
}

func (s *service) Receiver(ctx context.Context, roomId domain.RoomId, userId domain.UserId) (<-chan domain.Message, error) {
	return s.pubsub.SubscribeToMessages(ctx, roomId, userId)
}
