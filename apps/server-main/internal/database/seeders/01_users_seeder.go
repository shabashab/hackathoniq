package seeders

import (
	"database/sql"
	"time"

	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/models"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/services"
	"gorm.io/gorm"
)

type UsersSeeder struct {
	authCryptoService *services.AuthCryptoService
}

func NewUsersSeeder(
	authCryptoService *services.AuthCryptoService,
) Seeder {
	return &UsersSeeder{
		authCryptoService: authCryptoService,
	}
}

func (UsersSeeder) Name() (_ string) {
	return "01_users_seeder"
}

func (seeder UsersSeeder) Execute(db *gorm.DB) (_ error) {
	passwordHash, err := seeder.authCryptoService.CreatePasswordHash("password")

	if err != nil {
		return err
	}

	users := []*models.User{
		{
			FirstName: "Admin",
			LastName:  "Adminovich",
			EmailVerifiedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			PasswordHash: passwordHash,
			Role:         models.RoleAdmin,
		},
		{
			FirstName: "User 1",
			LastName:  "Userovich",
			EmailVerifiedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			PasswordHash: passwordHash,
			Role:         models.RoleUser,
		},
		{
			FirstName: "User 2",
			LastName:  "Userovich",
			EmailVerifiedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			PasswordHash: passwordHash,
			Role:         models.RoleUser,
		},
	}

	var result = db.Create(&users)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
