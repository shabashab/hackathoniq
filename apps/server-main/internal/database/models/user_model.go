package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	gorm.Model

	FirstName string
	LastName  string

	EmailVerifiedAt sql.NullTime

	PasswordHash string

	Role UserRole `gorm:"default:user"`
}
