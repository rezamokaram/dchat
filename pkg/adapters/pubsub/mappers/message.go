package mappers

import (
	"github.com/google/uuid"
	"github.com/rezamokaram/dchat/internal/chat/domain"
	"github.com/rezamokaram/dchat/pkg/adapters/pubsub/types"
)

func MessageSubscriber2Domain(msg types.Message) (*domain.Message, error) {
	domainMessage := domain.Message{
		Content: msg.Content,
		Filled:  true,
	}
	var err error
	userId, err := uuid.Parse(msg.UserId)
	if err != nil {
		return nil, err
	}

	roomId, err := uuid.Parse(msg.RoomId)
	if err != nil {
		return nil, err
	}

	domainMessage.UserId = domain.UserId(userId)
	domainMessage.RoomId = domain.RoomId(roomId)
	return &domainMessage, nil
}

func MessageDomain2Publisher(domainMsg domain.Message) (*types.Message, error) {
	return &types.Message{
		UserId:  domainMsg.UserId.ToString(),
		RoomId:  domainMsg.RoomId.ToString(),
		Content: domainMsg.Content,
	}, nil
}
