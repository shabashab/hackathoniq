package main

import (
	"github.com/shabashab/chattin/apps/chat-server/src/api"
	"github.com/shabashab/chattin/apps/chat-server/src/config"
	"github.com/shabashab/chattin/apps/chat-server/src/database"
	"github.com/shabashab/chattin/apps/chat-server/src/services"

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