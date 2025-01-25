package chat

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/RezaMokaram/chapp/api/pb"
	"github.com/RezaMokaram/chapp/internal/chat/domain"
	"github.com/google/uuid"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/internal/status"
)

func registerRecover(cancel context.CancelFunc) {
	if r := recover(); r != nil {
		log.Printf("PANIC CHAT LOOP: %v\n%s", r, string(debug.Stack()))
		cancel()
	}
}

func ChatRequest2DomainMessage(req *pb.ChatStreamRequest) (*domain.Message, error) {
	resp := domain.Message{
		Content: req.Content,
		Filled:  req.Filled,
	}
	roomId, err := uuid.Parse(req.RoomId)
	if err != nil {
		return nil, err
	}
	resp.RoomId = domain.RoomId(roomId)
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	resp.UserId = domain.UserId(userId)

	return &resp, err
}

func DomainMessage2ChatResponse(m domain.Message, remain uint64) *pb.ChatStreamResponse {
	return &pb.ChatStreamResponse{
		RoomId:  m.RoomId.ToString(),
		UserId:  m.UserId.ToString(),
		Content: m.Content,
		Filled:  true,
		Remain:  remain,
		Error:   "",
	}
}
