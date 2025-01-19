package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/RezaMokaram/chapp/api/pb"
	"github.com/RezaMokaram/chapp/app/room"
	"github.com/RezaMokaram/chapp/config"
	"google.golang.org/grpc"
)

func Run(appContainer room.RoomApp, cfg config.RoomConfig) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Room.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			contextUnaryInterceptor,
			setTransactionUnaryInterceptor(appContainer.DB()),
		),
	)

	pb.RegisterRoomServiceServer(s, &roomServer{})

	return s.Serve(lis)
}