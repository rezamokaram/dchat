package config

type ChatConfig struct {
	Chat CHAT `env-required:"false" json:"chat"`
	Nats NATS `env-required:"false" json:"nats"`
}

func (ChatConfig) configSignature() {}

type CHAT struct {
	Name         string `env-required:"false" json:"name" env:"CHAPP_CHAT_APP_NAME"`
	Version      string `env-required:"false" json:"version" env:"CHAPP_CHAT_APP_VERSION"`
	Host         string `env-required:"false" json:"host" env:"CHAPP_CHAT_APP_HOST"`
	Port         uint   `env-required:"false" json:"port" env:"CHAPP_CHAT_APP_PORT"`
	Phost 	string `env-required:"false" json:"phost" env:"CHAPP_CHAT_APP_PRESENCE_HOST"`
	Pport 	string `env-required:"false" json:"pport" env:"CHAPP_CHAT_APP_PRESENCE_PORT"`
}
