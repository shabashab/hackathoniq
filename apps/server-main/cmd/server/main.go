package main

import (
	"github.com/shabashab/hackathoniq/apps/server-main/internal/api"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/config"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/services"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		api.Module,
		services.Module,
		config.Module,
		database.Module,

		fx.Provide(
			zap.NewProduction,
		),
	).Run()
}
