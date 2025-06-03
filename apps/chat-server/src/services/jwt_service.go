package services

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shabashab/chattin/apps/chat-server/src/config/configs"
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
)

type JwtService struct {
	authConfig *configs.AuthConfig
	usersService *UsersService
}

func NewJwtService(authConfig *configs.AuthConfig, usersService *UsersService) (*JwtService) {
	return &JwtService{
		authConfig: authConfig,
		usersService: usersService,
	}
}

func (s JwtService) ValidateAndParseJwtToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.authConfig.JWTPrivateKey), nil
	})

	if err != nil {
		return nil, err
	}

	sub, err := token.Claims.GetSubject()

	if err != nil {
		return nil, err
	}

	subNumber, err := strconv.Atoi(sub)

	if err != nil {
		return nil, err
	}

	user, err := s.usersService.FindUserById(uint(subNumber))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s JwtService) CreateTokenForUserId(userId uint) (string, error) {
	user, err := s.usersService.FindUserById(userId)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": strconv.FormatUint(uint64(user.ID), 10),
		},
	)

	signedString, err := token.SignedString([]byte(s.authConfig.JWTPrivateKey))

	if err != nil {
		return "", err
	}

	return signedString, nil
}