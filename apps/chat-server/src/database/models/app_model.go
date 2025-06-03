package models

import "gorm.io/gorm"

type App struct {
	gorm.Model

	Name string
	AppKey string

	Users	[]User
}