package config

type ChatConfig struct {
	Chat CHAT `env-required:"false" json:"chat"`
	Nats NATS `env-required:"false" json:"nats"`
}

func (ChatConfig) configSignature() {}

type CHAT struct {
	Name    string `env-required:"false" json:"name" env:"dchat_CHAT_APP_NAME"`
	Version string `env-required:"false" json:"version" env:"dchat_CHAT_APP_VERSION"`
	Host    string `env-required:"false" json:"host" env:"dchat_CHAT_APP_HOST"`
	Port    uint   `env-required:"false" json:"port" env:"dchat_CHAT_APP_PORT"`
	Phost   string `env-required:"false" json:"phost" env:"dchat_CHAT_APP_PRESENCE_HOST"`
	Pport   string `env-required:"false" json:"pport" env:"dchat_CHAT_APP_PRESENCE_PORT"`
}
