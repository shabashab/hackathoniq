package configs

import "github.com/spf13/viper"

type AuthConfig struct {
	JWTPrivateKey string
}

func NewAuthConfig(v *viper.Viper) (*AuthConfig) {
	return &AuthConfig{
		JWTPrivateKey: v.GetString("AUTH_JWT_PRIVATE_KEY"),
	}
}