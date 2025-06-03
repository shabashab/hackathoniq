package main

import (
	"context"

	"github.com/shabashab/hackathoniq/apps/server-main/internal/config"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/config/configs"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MigrateParams struct {
	fx.In

	Lifecycle  fx.Lifecycle
	DB         *gorm.DB
	Models     []models.Model `group:"models"`
	DBConfig   *configs.DBConfig
	Logger     *zap.Logger
	Shutdowner fx.Shutdowner
}

func main() {
	fx.New(
		config.Module,
		database.Module,

		fx.Provide(
			zap.NewDevelopment,
		),
		fx.Invoke(autoMigrate),
	).Run()
}

func autoMigrate(p MigrateParams) {
	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, model := range p.Models {
				err := p.DB.AutoMigrate(model)

				if err != nil {
					return err
				}
			}

			p.Shutdowner.Shutdown()

			return nil
		},
	})
}
