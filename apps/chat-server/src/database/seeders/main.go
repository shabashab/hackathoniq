package seeders

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Seeder interface {
	Name() string
	Execute(db *gorm.DB) error
}

var Module = fx.Module("database.seeders",
	provideSeeder(NewAppsSeeder),
	provideSeeder(NewUsersSeeder),
)

func provideSeeder(seederFactory any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			seederFactory,
			fx.As(new(Seeder)),
			fx.ResultTags(`group:"seeders"`),
		),
	)
}