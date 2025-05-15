package interceptors

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/rezamokaram/dchat/api/handler/common"
	"github.com/rezamokaram/dchat/api/service"
	"github.com/rezamokaram/dchat/app/room"
	appContext "github.com/rezamokaram/dchat/pkg/context"
	"github.com/rezamokaram/dchat/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SvcContextKey = "svc"
)

func ContextUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	appCtx := appContext.NewAppContext(context.Background(), appContext.WithLogger(logger.NewLogger()))

	return handler(appCtx, req)
}

func SetRoomServiceGetterUnaryInterceptor(svc common.ServiceGetter[*service.RoomService]) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		withServiceCtx := context.WithValue(ctx, SvcContextKey, svc)
		return handler(withServiceCtx, req)
	}
}

func SetPresenceServiceGetterUnaryInterceptor(svc common.ServiceGetter[*service.PresenceService]) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		withServiceCtx := context.WithValue(ctx, SvcContextKey, svc)
		return handler(withServiceCtx, req)
	}
}

func SetTransactionUnaryInterceptor(appContainer room.RoomApp) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		tx := appContainer.DB().Begin()

		appContext.SetDB(ctx, tx, true)

		resp, err := handler(ctx, req)

		if err != nil {
			return nil, appContext.Rollback(ctx)
		}

		if err := appContext.CommitOrRollback(ctx, true); err != nil {
			return nil, err
		}
		return resp, err
	}
}

func LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		log.Printf("Received request: %v - Error occurred: %v", req, err)
	} else {
		log.Printf("Received request: %v - Sent response: %v", req, resp)
	}

	return resp, err
}

func PanicRecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("PANIC: %v\n%s", r, string(debug.Stack()))

			err = status.Errorf(codes.Internal, "internal server error")
		}
	}()

	return handler(ctx, req)
}
