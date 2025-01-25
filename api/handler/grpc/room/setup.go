package room

import (
	"fmt"
	"log"
	"net"

	"github.com/RezaMokaram/chapp/api/handler/grpc/interceptors"
	"github.com/RezaMokaram/chapp/api/pb"
	"github.com/RezaMokaram/chapp/app/room"
	"github.com/RezaMokaram/chapp/config"
	"google.golang.org/grpc"
)

func Run(appContainer room.RoomApp, cfg config.RoomConfig) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Room.Host, cfg.Room.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.ContextUnaryInterceptor,
			interceptors.SetTransactionUnaryInterceptor(appContainer),
			interceptors.SetRoomServiceGetterUnaryInterceptor(roomServiceGetter(appContainer)),
			interceptors.LoggingUnaryInterceptor,
			interceptors.PanicRecoveryInterceptor,
		),
	)

	pb.RegisterRoomServiceServer(s, &roomServer{
		appContainer: appContainer,
	})

	log.Printf("gRPC server is starting on %v:%v\n", cfg.Room.Host, cfg.Room.Port)
	return s.Serve(lis)
}
