package chat

import (
	"context"

	"github.com/RezaMokaram/chapp/api/pb"
	"github.com/RezaMokaram/chapp/app/chat"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	appContainer chat.ChatApp
}

func (s *chatServer) StreamMessages(stream pb.ChatService_StreamMessagesServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer registerRecover(cancel) // gracefully shutdown
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	domainMessage, err := ChatRequest2DomainMessage(req)
	if err != nil {
		return err
	}

	chatService := s.appContainer.ChatService(ctx)
	ch, err := chatService.Receiver(ctx, domainMessage.RoomId, domainMessage.UserId)
	if err != nil {
		return err
	}

	for {
		req, err = stream.Recv()
		if err != nil {
			panic(err)
		}

		domainMessage, err := ChatRequest2DomainMessage(req)
		if err != nil {
			err = stream.Send(&pb.ChatStreamResponse{
				Error:  err.Error(),
				Remain: uint64(len(ch)),
			})
			if err != nil {
				panic(err)
			}
			continue
		}

		if domainMessage.Filled {
			err = chatService.Send(ctx, *domainMessage)
			if err != nil {
				err = stream.Send(&pb.ChatStreamResponse{
					Error:  err.Error(),
					Remain: uint64(len(ch)),
				})
				if err != nil {
					panic(err)
				}
				continue
			}
		}

		if len(ch) > 0 {

			domainMsg := <-ch
			cnt := uint64(len(ch))
			if domainMessage.Filled {
				cnt++
			}

			err = stream.Send(DomainMessage2ChatResponse(domainMsg, cnt))
			if err != nil {
				panic(err)
			}
			continue
		} else {
			cnt := uint64(len(ch))
			if domainMessage.Filled {
				cnt++
			}

			err = stream.Send(&pb.ChatStreamResponse{
				Remain: cnt,
				Filled: false,
			})
			if err != nil {
				panic(err)
			}
			continue
		}
		break
	}

	cancel()
	return nil
}
