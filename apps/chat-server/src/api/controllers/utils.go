package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
)

func GetRequestAuthenticatedUserOrPanic(ctx *gin.Context) (*models.User) {
	possibleUser := ctx.MustGet("authenticated_user")

	user, ok := possibleUser.(*models.User)

	if !ok {
		// Shouldn't happen in normal situations
		panic("Can't convert authenticated_user to models.User")
	}

	return user
}