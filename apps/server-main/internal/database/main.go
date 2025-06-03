package database

import (
	"github.com/shabashab/hackathoniq/apps/server-main/internal/config/configs"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/models"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/repositories"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/seeders"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Module("database",
	fx.Provide(
		newDatabase,
	),

	models.Module,
	seeders.Module,
	repositories.Module,
)

func newDatabase(dbConfig *configs.DBConfig) (*gorm.DB, error) {
	pg := postgres.Open(dbConfig.ConnectionUrl)

	db, err := gorm.Open(pg)

	if err != nil {
		return nil, err
	}

	return db, nil
}
