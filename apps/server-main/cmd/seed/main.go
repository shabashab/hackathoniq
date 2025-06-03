package main

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/shabashab/hackathoniq/apps/server-main/internal/config"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/config/configs"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/models"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/seeders"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SeedParams struct {
	fx.In

	Lifecycle  fx.Lifecycle
	DB         *gorm.DB
	Seeders    []seeders.Seeder `group:"seeders"`
	DBConfig   *configs.DBConfig
	Logger     *zap.Logger
	Shutdowner fx.Shutdowner
}

func main() {
	fx.New(
		services.Module,
		config.Module,
		database.Module,

		fx.Provide(
			zap.NewDevelopment,
		),
		fx.Invoke(seed),
	).Run()
}

func seed(p SeedParams) {
	sort.Slice(p.Seeders, func(i, j int) bool {
		strI := p.Seeders[i].Name()
		strJ := p.Seeders[j].Name()

		return strings.Compare(strI, strJ) < 0
	})

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := p.DB.Transaction(func(tx *gorm.DB) error {
				for _, seeder := range p.Seeders {
					name := seeder.Name()

					seederExecuted, err := checkSeederExecuted(tx, name)

					if err != nil {
						p.Logger.Error("Failed to get if seeder has been executed", zap.Error(err))
						return err
					}

					if seederExecuted {
						p.Logger.Info("Skipping seeder as it has already been executed", zap.String("seeder_name", name))
						continue
					}

					p.Logger.Info("Executing seeder", zap.String("seeder_name", name))

					err = seeder.Execute(tx)

					if err != nil {
						p.Logger.Error("Seeder failed with error", zap.Error(err))
						return err
					}

					seed := &models.Seed{
						ExecutedAt: time.Now(),
						ID:         name,
					}

					result := tx.Create(&seed)

					if result.Error != nil {
						p.Logger.Info("Seeder executed successfully, but could not be saved", zap.Error(result.Error))
						return err
					}

					p.Logger.Info("Seeder executed successfully", zap.String("seeder_name", name))
				}

				p.Shutdowner.Shutdown()

				return nil
			})

			return err
		},
	})
}

func checkSeederExecuted(db *gorm.DB, seederName string) (bool, error) {
	seed := &models.Seed{}

	result := db.Where("id = ?", seederName).Limit(1).Find(&seed)

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}
