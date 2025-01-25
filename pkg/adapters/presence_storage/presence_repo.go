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
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	ErrRoomNotFound      = errors.New("room not found")
	ErrUserNotFound      = errors.New("user not found")
	ErrTransactionFailed = errors.New("transaction failed")
)

type presenceRepo struct {
	etcd *client.Client
	ttl  int64
	session *concurrency.Session
}

func NewPresenceRepo(etcd *client.Client, ttl int64) port.Repo {
	session, err := concurrency.NewSession(etcd)
	if err != nil {
		log.Println("can not create the session: %v", err)
	}

	// Create a mutex for the lock
	// mutex := concurrency.NewMutex(session, "/my-lock")

	return &presenceRepo{
		etcd: etcd,
		ttl:  ttl,
		session: session,
	}
}

func (pr *presenceRepo) SetUserPresence(ctx context.Context, userDomain domain.User) error {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mutex, err := pr.lock(ctxWithTimeOut, userDomain.RoomID.ToString())
	if err != nil {
		return err
	}
	defer pr.unlock(ctxWithTimeOut, mutex)

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
	mutex, err := pr.lock(ctxWithTimeOut, userDomainId.ToString())
	if err != nil {
		return err
	}
	defer pr.unlock(ctxWithTimeOut, mutex)

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
	mutex, err := pr.lock(ctxWithTimeOut, filter.ID.ToString())
	if err != nil {
		return nil, err
	}
	defer pr.unlock(ctxWithTimeOut, mutex)

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
	mutex, err := pr.lock(ctxWithTimeOut, filter.ID.ToString())
	if err != nil {
		return nil, err
	}
	defer pr.unlock(ctxWithTimeOut, mutex)

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

func (pr *presenceRepo) lock(ctx context.Context, key string) (*concurrency.Mutex, error) {
	mutex := concurrency.NewMutex(pr.session, key)
	if err := mutex.Lock(ctx); err != nil {
		return nil, err
	}
	return mutex, nil
}

func (pr *presenceRepo) unlock(ctx context.Context, mutex *concurrency.Mutex) {
	if err := mutex.Unlock(ctx); err != nil {
		log.Printf("error in unlocking th lock: %v", err)
	}
}