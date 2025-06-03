package services

import (
	"golang.org/x/crypto/bcrypt"
)

type AuthCryptoService struct{}

const bcryptHashingCost = 10

func NewAuthCryptoService() *AuthCryptoService {
	return &AuthCryptoService{}
}

func (s *AuthCryptoService) CreatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptHashingCost)
	return string(bytes), err
}

func (s *AuthCryptoService) VerifyPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
