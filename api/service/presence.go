package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rezamokaram/dchat/api/pb"
	"github.com/rezamokaram/dchat/internal/presence/domain"
	presencePort "github.com/rezamokaram/dchat/internal/presence/port"
)

type PresenceService struct {
	presenceSvc presencePort.Service
}

func NewPresenceService(
	presenceSvc presencePort.Service,
) *PresenceService {
	return &PresenceService{
		presenceSvc: presenceSvc,
	}
}

func (pr *PresenceService) GetRoomPresenceData(ctx context.Context, in *pb.GetRoomPresenceDataRequest) (*pb.GetRoomPresenceDataResponse, error) {
	roomId, err := uuid.Parse(in.RoomId)
	if err != nil {
		return nil, err
	}

	resp, err := pr.presenceSvc.GetRoomByFilter(ctx, domain.RoomFilter{
		ID: domain.RoomId(roomId),
	})
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("there is no room with this filter")
	}

	users := make([]*pb.UserPresenceData, 0)
	if resp.Users != nil {
		for _, usr := range resp.Users {
			users = append(users, &pb.UserPresenceData{
				UserId:    usr.ID.ToString(),
				RoomId:    usr.RoomID.ToString(),
				Status:    uint32(usr.Status),
				UpdatedAt: usr.UpdatedAt.String(),
			})
		}
	}
	return &pb.GetRoomPresenceDataResponse{
		Room: &pb.RoomData{
			RoomId: resp.ID.ToString(),
			Users:  users,
		},
	}, nil
}

func (pr *PresenceService) SetUserPresenceData(ctx context.Context, in *pb.SetUserPresenceRequest) (*pb.SetUserPresenceResponse, error) {
	userId, err := uuid.Parse(in.User.UserId)
	if err != nil {
		return nil, err
	}

	roomId, err := uuid.Parse(in.User.RoomId)
	if err != nil {
		return nil, err
	}

	err = pr.presenceSvc.SetUserPresence(ctx, domain.User{
		ID:        domain.UserId(userId),
		RoomID:    domain.RoomId(roomId),
		Status:    1,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.SetUserPresenceResponse{
		Success: true,
	}, nil
}

func (pr *PresenceService) DeleteUserPresenceData(ctx context.Context, in *pb.DeleteUserPresenceRequest) (*pb.DeleteUserPresenceResponse, error) {
	userId, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, err
	}

	err = pr.presenceSvc.DeleteUserPresence(ctx, domain.UserId(userId))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserPresenceResponse{
		Success: true,
	}, nil
}

func (pr *PresenceService) GetUserPresence(ctx context.Context, in *pb.GetUserPresenceRequest) (*pb.GetUserPresenceResponse, error) {
	userId, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, err
	}

	users, err := pr.presenceSvc.GetUsersByFilter(ctx, domain.UserFilter{ID: domain.UserId(userId)})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no user found")
	}

	return &pb.GetUserPresenceResponse{
		User: &pb.UserPresenceData{
			UserId:    users[0].ID.ToString(),
			RoomId:    users[0].RoomID.ToString(),
			Status:    uint32(users[0].Status),
			UpdatedAt: users[0].UpdatedAt.String(),
		},
	}, nil
}

func (pr *PresenceService) GetUsersPresence(ctx context.Context, in *pb.GetUsersPresenceRequest) (*pb.GetUsersPresenceResponse, error) {
	domainUsers, err := pr.presenceSvc.GetUsersByFilter(ctx, domain.UserFilter{})
	if err != nil {
		return nil, err
	}

	if len(domainUsers) == 0 {
		return nil, errors.New("no user found")
	}
	users := make([]*pb.UserPresenceData, 0)
	for _, usr := range domainUsers {
		users = append(users, &pb.UserPresenceData{
			UserId:    usr.ID.ToString(),
			RoomId:    usr.RoomID.ToString(),
			Status:    uint32(usr.Status),
			UpdatedAt: usr.UpdatedAt.String(),
		})
	}
	return &pb.GetUsersPresenceResponse{
		Users: users,
	}, nil
}
