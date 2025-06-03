package configs

import "github.com/spf13/viper"

type ApiConfig struct {
	Host string
}

func NewApiConfig(v *viper.Viper) (*ApiConfig) {
	return &ApiConfig{
		Host: v.GetString("API_HOST"),
	}
}