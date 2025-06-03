package configs

import "github.com/spf13/viper"

type DBConfig struct {
	ConnectionUrl string

	AutoMigrate bool
	AutoSeed bool
}

func NewDbConfig(v *viper.Viper) (*DBConfig) {
	v.SetDefault("DATABASE_AUTO_MIGRATE", false)
	v.SetDefault("DATABASE_AUTO_SEED", false)

	return &DBConfig{
		ConnectionUrl: v.GetString("DATABASE_URL"),
		AutoMigrate: v.GetBool("DATABASE_AUTO_MIGRATE"),
		AutoSeed: v.GetBool("DATABASE_AUTO_SEED"),
	}
}