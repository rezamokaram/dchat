package config

type POSTGRES struct {
	DB       string `env-required:"false" json:"db" env:"dchat_POSTGRES_DB"`
	User     string `env-required:"false" json:"user" env:"dchat_POSTGRES_USER"`
	Password string `env-required:"false" json:"password" env:"dchat_POSTGRES_PASSWORD"`
	Host     string `env-required:"false" json:"host" env:"dchat_POSTGRES_HOST"`
	Port     uint   `env-required:"false" json:"port" env:"dchat_POSTGRES_PORT"`
	SSLMode  string `env-required:"false" json:"sslmode" env:"dchat_POSTGRES_SSLMODE"`
	Timezone string `env-required:"false" json:"timezone" env:"dchat_POSTGRES_TIMEZONE"`
	Schema   string `env-required:"false" json:"schema" env:"dchat_POSTGRES_SCHEMA"`
}

type REDIS struct {
	Host string `env-required:"false" json:"host" env:"dchat_REDIS_HOST"`
	Port uint   `env-required:"false" json:"port" env:"dchat_REDIS_PORT"`
}

type ETCD struct {
	Hosts []string `env-required:"false" json:"hosts" env:"dchat_ETCD_HOST"`
	TTL   int64    `env-required:"false" json:"ttl" env:"dchat_ETCD_TTL"`
}

type NATS struct {
	Host string `env-required:"false" json:"host" env:"dchat_NATS_HOST"`
}

type SCYLLA struct {
	Hosts            []string `env-required:"false" json:"hosts" env:"dchat_SCYLLA_HOSTS"`
	Keyspace         string   `env-required:"false" json:"keyspace" env:"dchat_SCYLLA_KEYSPACE"`
	ConsistencyLevel string   `env-required:"false" json:"consistency_level" env:"dchat_SCYLLA_CONSISTENCY_LEVEL"`
	ProtoVersion     int      `env-required:"false" json:"proto_version" env:"dchat_SCYLLA_PROTO_VERSION"`
	ConnectTimeout   int      `env-required:"false" json:"connect_timeout" env:"dchat_SCYLLA_CONNECTION_TIMEOUT"`
	Timeout          int      `env-required:"false" json:"timeout" env:"dchat_SCYLLA_TIMEOUT"`
}
