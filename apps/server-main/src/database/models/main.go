package models

import (
	"go.uber.org/fx"
)

var Module = fx.Module("database.models",
	provideModel(&User{}),
	provideModel(&App{}),
	provideModel(&Seed{}),
)

type Model interface{}

func provideModel(model Model) fx.Option {
	return fx.Provide(
		asModel(model),
	)
}

func asModel(model Model) any {
	newModel := func() any {
		return model
	}

	return fx.Annotate(
		newModel,
		fx.As(new(Model)),
		fx.ResultTags(`group:"models"`),
	)
}