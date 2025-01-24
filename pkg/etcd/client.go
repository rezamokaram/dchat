package etcd

import (
	"time"

	client "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient(endpoints []string) (*client.Client, error) {
	config := client.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}

	c, err := client.New(config)
	if err != nil {
		return nil, err
	}

	return c, nil
}
