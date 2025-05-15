package pubsub

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/rezamokaram/dchat/internal/chat/domain"
	"github.com/rezamokaram/dchat/internal/chat/port"
	"github.com/rezamokaram/dchat/pkg/adapters/pubsub/mappers"
	"github.com/rezamokaram/dchat/pkg/adapters/pubsub/types"
)

var (
	ErrNoSubscription = errors.New("pubsub do not start any subscription")
)

type pubsub struct {
	nc  *nats.Conn
	sub *nats.Subscription
}

// this type of subscribing is not a good way to do it
// but in my system design the source of the message is a another storage (scylla db in message service)

// another good idea is to design this adapter in non-depend for messaging system
// for example using watermill or using a good interface

func NewPubSub(nc *nats.Conn) port.Pubsub {
	return &pubsub{
		nc:  nc,
		sub: nil,
	}
}

func (s *pubsub) SubscribeToMessages(ctx context.Context, roomId domain.RoomId, userId domain.UserId) (<-chan domain.Message, error) {
	msgs := make(chan domain.Message, 1000)
	sub, err := s.nc.Subscribe(getSubjectName(roomId.ToString()), func(m *nats.Msg) {
		var receivedMsg types.Message
		if err := json.Unmarshal(m.Data, &receivedMsg); err != nil {
			log.Printf("MESSAGE SUBSCRIBER ERROR: error while unmarshalling message: %v", err)
			return
		}

		domainMessage, err := mappers.MessageSubscriber2Domain(receivedMsg)
		if err != nil {
			log.Printf("MESSAGE SUBSCRIBER ERROR: error while mapping message: %v", err)
			return
		}

		// this is a policy :(:): FUN TIME :)
		if len(msgs) > 999 {
			<-msgs
		}
		msgs <- *domainMessage
	})

	if err != nil {
		return nil, err
	}
	s.sub = sub
	return msgs, nil
}

func (s *pubsub) UnSubscribe(ctx context.Context) error {
	if s.sub == nil {
		return ErrNoSubscription
	}
	return s.sub.Unsubscribe()
}

func (s *pubsub) PublishMessage(
	ctx context.Context,
	message domain.Message,
) error {
	msg, err := mappers.MessageDomain2Publisher(message)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	subject := getSubjectName(message.RoomId.ToString())
	err = s.nc.Publish(subject, bytes)
	if err != nil {
		return err
	}

	return nil
}
