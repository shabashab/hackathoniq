package repositories

import (
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) (*UsersRepository) {
	return &UsersRepository{
		db: db,
	}
}

func (r UsersRepository) FindUserById(id uint) (*models.User, error) {
	user := &models.User{}

	result := r.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
