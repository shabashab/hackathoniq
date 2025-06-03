package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shabashab/chattin/apps/chat-server/src/api/dtos"
	"github.com/shabashab/chattin/apps/chat-server/src/services"
)

type AuthController struct {
	jwtService *services.JwtService
}

type authenticateDebugBody struct {
	UserId uint `json:"userId"`
}

func NewAuthController(jwtService *services.JwtService) *AuthController {
	return &AuthController{
		jwtService: jwtService,
	}
}

// GetCurrentUser retrieves the currently authenticated user.
//
// @Summary Get current authenticated user
// @Description Retrieves information about the currently logged-in user
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dtos.IamDto
// @Router /auth/iam [get]
func (c AuthController) GetCurrentUser(ctx *gin.Context) {
	authenticatedUser := GetRequestAuthenticatedUserOrPanic(ctx)

	ctx.JSON(http.StatusOK, dtos.NewIamDto(authenticatedUser))
}

// DebugLogin allows debugging authentication by generating a JWT token.
//
// @Summary Debug login
// @Description Generates a JWT token for a given user ID (only for debugging purposes)
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body authenticateDebugBody true "User ID for authentication"
// @Success 200 {object} dtos.DebugLoginDto
// @Router /auth/debug/login [post]
func (c AuthController) DebugLogin(ctx *gin.Context) {
	body := &authenticateDebugBody{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := c.jwtService.CreateTokenForUserId(body.UserId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewDebugLoginDto(token))
}
