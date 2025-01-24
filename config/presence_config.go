package config

type PresenceConfig struct {
	Presence PRESENCE `env-required:"false" json:"presence"`
	Etcd     ETCD     `env-required:"false" json:"etcd"`
}

func (PresenceConfig) configSignature() {}

type PRESENCE struct {
	Name    string `env-required:"false" json:"name" env:"CHAPP_PRESENCE_APP_NAME"`
	Version string `env-required:"false" json:"version" env:"CHAPP_PRESENCE_APP_VERSION"`
	Host    string `env-required:"false" json:"host" env:"CHAPP_PRESENCE_APP_HOST"`
	Port    uint   `env-required:"false" json:"port" env:"CHAPP_PRESENCE_APP_PORT"`
}
