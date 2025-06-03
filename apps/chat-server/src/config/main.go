package config

import (
	"github.com/shabashab/chattin/apps/chat-server/src/config/configs"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(
		NewViper,
	),

	configs.Module,
)

func NewViper() (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile("config/.env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return v, nil
}