package config

type RoomConfig struct {
	Room     ROOM     `env-required:"false" json:"room"`
	Postgres POSTGRES `env-required:"false" json:"postgres"`
	Redis    REDIS    `env-required:"false" json:"redis"`
}

func (RoomConfig) configSignature() {}

type ROOM struct {
	Name    string `env-required:"false" json:"name" env:"dchat_APP_NAME"`
	Version string `env-required:"false" json:"version" env:"dchat_APP_VERSION"`
	Host    string `env-required:"false" json:"host" env:"dchat_APP_HOST"`
	Port    uint   `env-required:"false" json:"port" env:"dchat_APP_PORT"`
}
