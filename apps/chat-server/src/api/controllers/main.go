package controllers

import "go.uber.org/fx"

var Module = fx.Module("api.controllers",
	fx.Provide(
		NewHealthController,
		NewAuthController,
	),
)