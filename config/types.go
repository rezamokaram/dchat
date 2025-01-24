package config

type POSTGRES struct {
	DB       string `env-required:"false" json:"db" env:"CHAPP_POSTGRES_DB"`
	User     string `env-required:"false" json:"user" env:"CHAPP_POSTGRES_USER"`
	Password string `env-required:"false" json:"password" env:"CHAPP_POSTGRES_PASSWORD"`
	Host     string `env-required:"false" json:"host" env:"CHAPP_POSTGRES_HOST"`
	Port     uint   `env-required:"false" json:"port" env:"CHAPP_POSTGRES_PORT"`
	SSLMode  string `env-required:"false" json:"sslmode" env:"CHAPP_POSTGRES_SSLMODE"`
	Timezone string `env-required:"false" json:"timezone" env:"CHAPP_POSTGRES_TIMEZONE"`
	Schema   string `env-required:"false" json:"schema" env:"CHAPP_POSTGRES_SCHEMA"`
}

type REDIS struct {
	Host string `env-required:"false" json:"host" env:"CHAPP_REDIS_HOST"`
	Port uint   `env-required:"false" json:"port" env:"CHAPP_REDIS_PORT"`
}

type ETCD struct {
	Hosts []string `env-required:"false" json:"hosts" env:"CHAPP_ETCD_HOST"`
	TTL int64 `env-required:"false" json:"ttl" env:"CHAPP_ETCD_TTL"`
}
