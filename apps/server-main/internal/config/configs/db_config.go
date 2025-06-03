package configs

import "github.com/spf13/viper"

type DBConfig struct {
	ConnectionUrl string
}

func NewDbConfig(v *viper.Viper) *DBConfig {
	return &DBConfig{
		ConnectionUrl: v.GetString("DATABASE_URL"),
	}
}
