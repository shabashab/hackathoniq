package dtos

import "github.com/shabashab/hackathoniq/apps/server-main/internal/database/models"

type IamDto struct {
	Id uint `json:"id"`
}

func NewIamDto(user *models.User) *IamDto {
	return &IamDto{
		Id: user.ID,
	}
}
