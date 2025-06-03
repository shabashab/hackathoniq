package dtos

import "github.com/shabashab/chattin/apps/chat-server/src/database/models"

type IamDto struct {
	Id uint `json:"id"`
}

func NewIamDto(user *models.User) (*IamDto) {
	return &IamDto{
		Id: user.ID,
	}
}