package scylla

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/rezamokaram/dchat/config"
	"github.com/scylladb/gocqlx/v2"
)

func NewScyllaClient(cfg config.SCYLLA) (*gocqlx.Session, error) {
	cluster := gocql.NewCluster(cfg.Hosts...)

	cluster.Keyspace = cfg.Keyspace
	cl, err := gocql.ParseConsistencyWrapper(cfg.ConsistencyLevel)
	if err != nil {
		return nil, err
	}

	cluster.Consistency = cl
	cluster.ProtoVersion = cfg.ProtoVersion
	cluster.ConnectTimeout = time.Second * time.Duration(cfg.ConnectTimeout)
	cluster.Timeout = time.Second * time.Duration(cfg.Timeout)

	// Optional settings
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy()) // todo
	cluster.DisableInitialHostLookup = true                                                           // todo

	// Create session
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	log.Println("Connected to ScyllaDB!")

	return &session, nil
}
