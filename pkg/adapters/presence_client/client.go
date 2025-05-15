package presence_client

import (
	"context"
	"time"

	"github.com/rezamokaram/dchat/api/pb"
	"github.com/rezamokaram/dchat/internal/chat/port"
	"google.golang.org/grpc"
)

type presenceClient struct {
	conn   *grpc.ClientConn
	client pb.PresenceServiceClient
}

func NewPresenceClient(presenceHost string) (port.PresenceClient, error) {
	conn, err := grpc.Dial(presenceHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := pb.NewPresenceServiceClient(conn)

	return &presenceClient{
		conn:   conn,
		client: client,
	}, nil
}

func (pc *presenceClient) SetUserPresence(ctx context.Context, roomId string, userId string) error {
	req := pb.SetUserPresenceRequest{
		User: &pb.UserPresenceData{
			UserId:    userId,
			RoomId:    roomId,
			Status:    1,
			UpdatedAt: time.Now().String(),
		},
	}
	_, err := pc.client.SetUserPresence(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
