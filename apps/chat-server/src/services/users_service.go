package services

import (
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
	"github.com/shabashab/chattin/apps/chat-server/src/database/repositories"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) (*UsersService) {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s UsersService) FindUserById(id uint) (*models.User, error) {
	return s.usersRepository.FindUserById(id)
}