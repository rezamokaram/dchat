package nats

import (
	"github.com/nats-io/nats.go"
)

func NewNatsClient(host string) (*nats.Conn, error) {
	nc, err := nats.Connect(host)
	if err != nil {
		return nil, err
	}
	return nc, nil
}
