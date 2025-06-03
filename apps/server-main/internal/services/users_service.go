package services

import (
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/models"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/database/repositories"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s UsersService) FindUserById(id uint) (*models.User, error) {
	return s.usersRepository.FindUserById(id)
}
