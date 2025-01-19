package grpc

import (
	"context"

	"github.com/RezaMokaram/chapp/api/pb"
)

type roomServer struct {
	pb.UnimplementedRoomServiceServer
}

func (s *roomServer) SignInUser(ctx context.Context, in *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	// todo log
	// log.Printf("Received: %v", in.GetName())
	return &pb.UserSignInResponse{}, nil
}

func (s *roomServer) SignUpUser(ctx context.Context, in *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	// todo log
	// log.Printf("Received: %v", in.GetName())
	return &pb.UserSignUpResponse{}, nil
}

func (s *roomServer) CreateRoom(ctx context.Context, in *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	// todo log
	// log.Printf("Received: %v", in.GetName())
	return &pb.CreateRoomResponse{}, nil
}
