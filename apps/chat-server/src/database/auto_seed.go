package database

import (
	"sort"
	"strings"
	"time"

	"github.com/shabashab/chattin/apps/chat-server/src/config/configs"
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
	"github.com/shabashab/chattin/apps/chat-server/src/database/seeders"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AutoSeedParams struct {
	fx.In

	DB *gorm.DB
	Seeders []seeders.Seeder `group:"seeders"`
	DBConfig *configs.DBConfig
	Logger *zap.Logger
}

func autoSeed(p AutoSeedParams) (error) {
	if(!p.DBConfig.AutoSeed) {
		p.Logger.Info("Auto seeding disabled, skipping")
		return nil
	}

	// Sorting lexicographically so we would have all the seeders executed in the correct order
	sort.Slice(p.Seeders, func(i, j int) bool {
		strI := p.Seeders[i].Name()
		strJ := p.Seeders[j].Name()

		return strings.Compare(strI, strJ) < 0
	})

	err := p.DB.Transaction(func(tx *gorm.DB) error {
		for _, seeder := range p.Seeders {
			name := seeder.Name()

			seederExecuted, err := checkSeederExecuted(tx, name)

			if(err != nil) {
				p.Logger.Error("Failed to get if seeder has been executed", zap.Error(err))
				return err
			}

			if(seederExecuted) {
				p.Logger.Info("Skipping seeder as it has already been executed", zap.String("seeder_name", name))
				continue
			}

			p.Logger.Info("Executing seeder", zap.String("seeder_name", name))

			err = seeder.Execute(tx)

			if(err != nil) {
				p.Logger.Error("Seeder failed with error", zap.Error(err))
				return err
			}

			seed := &models.Seed{
				ExecutedAt: time.Now(),
				ID: name,
			}

			result := tx.Create(&seed)

			if(result.Error != nil) {
				p.Logger.Info("Seeder executed successfully, but could not be saved", zap.Error(result.Error))
				return err
			}

			p.Logger.Info("Seeder executed successfully", zap.String("seeder_name", name))
		}

		return nil
	})

	return err
}

func checkSeederExecuted(db *gorm.DB, seederName string) (bool, error) {
	seed := &models.Seed{}

	result := db.Where("id = ?", seederName).Limit(1).Find(&seed)

	if(result.Error != nil) {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}