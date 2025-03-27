package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"API3/registro/src/application"
	"API3/registro/src/domain/entities"
)

type UserController struct {
	UseCase *user_usecase.UserUseCase
}

func NewUserController(uc *user_usecase.UserUseCase) *UserController {
	return &UserController{UseCase: uc}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := c.UseCase.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdUser)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	user, err := c.UseCase.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.UseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := c.UseCase.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	if err := c.UseCase.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
