package in_memory_kv_storage

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/RezaMokaram/chapp/internal/presence/domain"
	"github.com/RezaMokaram/chapp/internal/presence/port"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_storage/mappers"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_storage/types"
	"github.com/google/uuid"
	client "go.etcd.io/etcd/client/v3"
)

type presenceRepo struct {
	etcd *client.Client
}

func NewPresenceRepo(etcd *client.Client) port.Repo {
	return &presenceRepo{
		etcd,
	}
}

func (pr *presenceRepo) SetUserPresence(ctx context.Context, userDomain domain.User) (error) {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	
	user:= mappers.UserDomain2Storage(userDomain)
	user.Status = 1
	user.UpdatedAt = time.Now()


	resp, err := pr.etcd.Get(ctxWithTimeOut, getRoomKey(user.RoomID))
	if err != nil {
		return err
	}

	room := types.Room {
		ID: user.RoomID,
		Users: make(map[uuid.UUID]types.User),
	}
	if resp.Count != 0 {
		json.Unmarshal(resp.Kvs[0].Value, &room)
	}

	room.Users[user.ID] = *user

	userVal, err := struct2String(user)
	if err != nil {
		return err
	}

	roomVal, err := struct2String(room)
	if err != nil {
		return err
	}

	txn := pr.etcd.Txn(ctxWithTimeOut)
	txn.Then(
		client.OpPut(getUserKey(user.ID), userVal), // TODO: lease
		client.OpPut(getRoomKey(room.ID), roomVal),
	)
	tnxResp, err := txn.Commit()
	if err != nil {
		return err
	}

	if !tnxResp.Succeeded {
		return errors.New("transaction failed")
	}

	return nil
}

func (pr *presenceRepo) DeleteUserPresence(ctx context.Context, userDomain domain.User) (error) {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	
	user:= mappers.UserDomain2Storage(userDomain)

	resp, err := pr.etcd.Get(ctxWithTimeOut, getRoomKey(user.RoomID))
	if err != nil {
		return err
	}

	room := types.Room {
		ID: user.RoomID,
		Users: make(map[uuid.UUID]types.User),
	}
	if resp.Count != 0 {
		json.Unmarshal(resp.Kvs[0].Value, &room)
	}

	delete(room.Users, user.ID)

	roomVal, err := struct2String(room)
	if err != nil {
		return err
	}

	txn := pr.etcd.Txn(ctxWithTimeOut)
	txn.Then(
		client.OpDelete(getUserKey(user.ID)),
		client.OpPut(getRoomKey(room.ID), roomVal),
	)
	tnxResp, err := txn.Commit()
	if err != nil {
		return err
	}

	if !tnxResp.Succeeded {
		return errors.New("transaction failed")
	}

	return nil
}

func (pr *presenceRepo) GetUsersByFilter(ctx context.Context, filter domain.UserFilter) ([]domain.User, error) {
	panic("implement me")
	// ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	// defer cancel()

	// resp, err := pr.etcd.Get(ctxWithTimeOut, getUserKey(uuid.UUID(filter.ID)))
	// if err != nil {
	// 	return nil, err
	// }

	// if resp.Count == 0 {
	// 	errors.New("user not found")
	// }

	// var user types.User
	// err = json.Unmarshal(resp.Kvs[0].Value, &user)
	// if err != nil {
	// 	return nil, err
	// }

	// return nil,nil//mappers.UserStorage2Domain(user) , nil
}

func (pr *presenceRepo) GetRoomByFilter(ctx context.Context, filter domain.RoomFilter) (*domain.Room, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	resp, err := pr.etcd.Get(ctxWithTimeOut, getRoomKey(uuid.UUID(filter.ID)))
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		errors.New("room not found")
	}

	var room types.Room
	err = json.Unmarshal(resp.Kvs[0].Value, &room)
	if err != nil {
		return nil, err
	}

	return mappers.RoomStorage2Domain(room) , nil
}

// private functions

// func (pr *presenceRepo) getUserCreateIfNotExist(ctx context.Context, txn client.Txn, userId uuid.UUID) (types.User, error) {
// 	panic("implement me")
// }

// func (pr *presenceRepo) getRoomCreateIfNotExist(ctx context.Context, txn client.Txn, roomId uuid.UUID) (types.Room, error) {
// 	panic("implement me")
// }