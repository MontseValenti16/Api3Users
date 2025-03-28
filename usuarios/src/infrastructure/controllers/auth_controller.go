package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"API3/usuarios/src/application"
)

type AuthController struct {
	UseCase *auth_usecase.AuthUseCase
}

func NewAuthController(uc *auth_usecase.AuthUseCase) *AuthController {
	return &AuthController{UseCase: uc}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.UseCase.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, LoginResponse{Token: token})
}
