package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shabashab/chattin/apps/chat-server/src/services"
)

type AuthMiddleware struct {
	jwtService *services.JwtService
}

func NewAuthMiddleware(jwtService *services.JwtService) (*AuthMiddleware) {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (middleware AuthMiddleware) Handler(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")

	if len(authorizationHeader) == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no_authorization_header_provided",
		})
		return
	}

	authHeaderStrings := strings.Split(authorizationHeader, " ")

	if len(authHeaderStrings) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid_authorization_header",
		})
		return
	}

	authType := authHeaderStrings[0]

	if authType != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid_authorization_header",
		})
		return
	}

	tokenString := authHeaderStrings[1]

	user, err := middleware.jwtService.ValidateAndParseJwtToken(tokenString)

	if (err != nil) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid_authorization_token",
		})
		return
	}

	ctx.Set("authenticated_user", user)

	ctx.Next()
}