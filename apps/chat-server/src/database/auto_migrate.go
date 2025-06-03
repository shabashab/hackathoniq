package database

import (
	"github.com/shabashab/chattin/apps/chat-server/src/config/configs"
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AutoMigrateParams struct {
	fx.In

	DB *gorm.DB
	Models []models.Model `group:"models"`
	DBConfig *configs.DBConfig
	Logger *zap.Logger
}

func autoMigrate(p AutoMigrateParams) (error) {
	if(!p.DBConfig.AutoMigrate) {
		p.Logger.Info("Auto migrating disabled, skipping")
		return nil
	}

	for _, model := range p.Models {
		err := p.DB.AutoMigrate(model)

		if(err != nil) {
			return err
		}
	}

	return nil
}