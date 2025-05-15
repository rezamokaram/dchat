package chat

import (
	"fmt"
	"log"
	"net"

	"github.com/rezamokaram/dchat/api/pb"
	"github.com/rezamokaram/dchat/app/chat"
	"github.com/rezamokaram/dchat/config"
	"google.golang.org/grpc"
)

func Run(appContainer chat.ChatApp, cfg config.ChatConfig) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Chat.Host, cfg.Chat.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &chatServer{
		appContainer: appContainer,
	})

	log.Printf("gRPC server is starting on %v:%v\n", cfg.Chat.Host, cfg.Chat.Port)
	return s.Serve(lis)
}
