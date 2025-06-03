package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	// Optional user's tag given from external system
	// Could be an external id or username
	Tag 	*string

	AppID uint
	App		App
}