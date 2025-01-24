package presence

import (
	"fmt"
	"log"
	"net"

	"github.com/RezaMokaram/chapp/api/handler/grpc/interceptors"
	"github.com/RezaMokaram/chapp/api/pb"
	"github.com/RezaMokaram/chapp/app/presence"
	"github.com/RezaMokaram/chapp/config"
	"google.golang.org/grpc"
)

func Run(appContainer presence.PresenceApp, cfg config.PresenceConfig) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Presence.Host, cfg.Presence.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.ContextUnaryInterceptor,
			interceptors.SetPresenceServiceGetterUnaryInterceptor(presenceServiceGetter(appContainer)),
			interceptors.LoggingUnaryInterceptor,
			interceptors.PanicRecoveryInterceptor,
		),
	)

	pb.RegisterPresenceServiceServer(s, &presenceServer{
		appContainer: appContainer,
	})

	log.Printf("gRPC server is starting on %v:%v\n", cfg.Presence.Host, cfg.Presence.Port)
	return s.Serve(lis)
}
