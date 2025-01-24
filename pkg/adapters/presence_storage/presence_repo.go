package in_memory_kv_storage

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/RezaMokaram/chapp/internal/presence/domain"
	"github.com/RezaMokaram/chapp/internal/presence/port"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_storage/mappers"
	"github.com/RezaMokaram/chapp/pkg/adapters/presence_storage/types"
	"github.com/google/uuid"
	client "go.etcd.io/etcd/client/v3"
)

var (
	ErrRoomNotFound      = errors.New("room not found")
	ErrUserNotFound      = errors.New("user not found")
	ErrTransactionFailed = errors.New("transaction failed")
)

type presenceRepo struct {
	etcd *client.Client
	ttl int64
}

func NewPresenceRepo(etcd *client.Client, ttl int64) port.Repo {
	return &presenceRepo{
		etcd: etcd,
		ttl: ttl,
	}
}

func (pr *presenceRepo) SetUserPresence(ctx context.Context, userDomain domain.User) error {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := mappers.UserDomain2Storage(userDomain)
	user.Status = 1
	user.UpdatedAt = time.Now()

	resp, err := pr.etcd.Get(ctxWithTimeOut, getRoomKey(user.RoomID))
	if err != nil {
		return err
	}

	room := types.Room{
		ID:    user.RoomID,
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

	lease, err := pr.etcd.Grant(context.Background(), pr.ttl)
	if err != nil {
		log.Fatalf("Failed to create lease: %v", err)
	}

	txn := pr.etcd.Txn(ctxWithTimeOut)
	txn.Then(
		client.OpPut(getUserKey(user.ID), userVal, client.WithLease(lease.ID)),
		client.OpPut(getRoomKey(room.ID), roomVal),
	)
	tnxResp, err := txn.Commit()
	if err != nil {
		return err
	}

	if !tnxResp.Succeeded {
		return ErrTransactionFailed
	}

	return nil
}

func (pr *presenceRepo) DeleteUserPresence(ctx context.Context, userDomainId domain.UserId) error {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := pr.etcd.Get(ctxWithTimeOut, getUserKey(uuid.UUID(userDomainId)))
	if err != nil {
		return err
	}
	if resp.Count == 0 {
		return ErrUserNotFound
	}
	var user types.User
	err = json.Unmarshal(resp.Kvs[0].Value, &user)
	if err != nil {
		return err
	}

	resp, err = pr.etcd.Get(ctxWithTimeOut, getRoomKey(user.RoomID))
	if err != nil {
		return err
	}

	room := types.Room{
		ID:    user.RoomID,
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
		return ErrTransactionFailed
	}

	return nil
}

func (pr *presenceRepo) GetUsersByFilter(ctx context.Context, filter domain.UserFilter) ([]domain.User, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	prefix := getUserKey(uuid.UUID(filter.ID))
	resp, err := pr.etcd.Get(ctxWithTimeOut, prefix, client.WithPrefix())
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, ErrUserNotFound
	}

	domainUsers := make([]domain.User, 0)
	for _, val := range resp.Kvs {
		var storageUser types.User
		err = json.Unmarshal(val.Value, &storageUser)
		if err != nil {
			log.Printf("error in unmarshal the etcd user data\n")
		}
		domainUsers = append(domainUsers, *mappers.UserStorage2Domain(storageUser))
	}
	return domainUsers, nil
}

func (pr *presenceRepo) GetRoomByFilter(ctx context.Context, filter domain.RoomFilter) (*domain.Room, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := pr.etcd.Get(ctxWithTimeOut, getRoomKey(uuid.UUID(filter.ID)))
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, ErrRoomNotFound
	}

	var room types.Room
	err = json.Unmarshal(resp.Kvs[0].Value, &room)
	if err != nil {
		return nil, err
	}
	return mappers.RoomStorage2Domain(room), nil
}