package seeders

import (
	"fmt"

	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
	"gorm.io/gorm"
)

type UsersSeeder struct {}

func NewUsersSeeder() (Seeder) {
	return &UsersSeeder{}
}

func (UsersSeeder) Name() (_ string) {
	return "02_users_seeder"
}

func (seeder UsersSeeder) Execute(db *gorm.DB) (_ error) {
	apps := []*models.App{}

	result := db.Find(&apps)

	if (result.Error != nil) {
		return result.Error
	}

	for _, app := range apps {
		users := []*models.User{
			{AppID: app.ID, Tag: seeder.createUserTag(app, 1)},
			{AppID: app.ID, Tag: seeder.createUserTag(app, 2)},
			{AppID: app.ID, Tag: seeder.createUserTag(app, 3)},
		}

		result = db.Create(&users)

		if (result.Error != nil) {
			return result.Error
		}
	}

	return nil
}

func (UsersSeeder) createUserTag(app *models.App, index int) (*string) {
	tag := fmt.Sprintf("%s-user-%d", app.Name, index)
	return &tag
}