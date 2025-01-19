package grpc

import (
	"context"

	appContext "github.com/RezaMokaram/chapp/pkg/context"
	"github.com/RezaMokaram/chapp/pkg/logger"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func contextUnaryInterceptor(
	ctx context.Context, 
	req interface{}, 
	info *grpc.UnaryServerInfo, 
	handler grpc.UnaryHandler,
	) (interface{}, error) {

	appCtx := appContext.NewAppContext(context.Background(), appContext.WithLogger(logger.NewLogger()))

	return handler(appCtx, req)
}

func setTransactionUnaryInterceptor(db *gorm.DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		tx := db.Begin()

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
