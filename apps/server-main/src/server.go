package main

import (
	"github.com/shabashab/hackathoniq/apps/server-main/src/api"
	"github.com/shabashab/hackathoniq/apps/server-main/src/config"
	"github.com/shabashab/hackathoniq/apps/server-main/src/database"
	"github.com/shabashab/hackathoniq/apps/server-main/src/services"

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
