package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"API3/usuarios/src/application"
	"API3/usuarios/src/domain/entities"
)

type AuthController struct {
	UseCase *auth_usecase.AuthUseCase
}

func NewAuthController(uc *auth_usecase.AuthUseCase) *AuthController {
	return &AuthController{UseCase: uc}
}

type LoginRequest struct {
	NombreUsuario    string `json:"nombre_usuario"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	IDUsuario      int    `json:"id_usuario"`
	NombreUsuario  string `json:"nombre_usuario"`
}

// RegisterRequest define los campos necesarios para registrar un usuario.
// Se espera que el front envíe el id_persona (ya obtenido tras registrar la persona),
// nombre_usuario, email y password.
type RegisterRequest struct {
	IDPersona     int    `json:"id_persona"`
	NombreUsuario string `json:"nombre_usuario"`
	Password      string `json:"password"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, userID, username, err := c.UseCase.Login(req.NombreUsuario, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Authorization", "Bearer "+token)
	
	ctx.JSON(http.StatusOK, LoginResponse{Token: token,
		IDUsuario: userID,
		NombreUsuario: username})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Convertir el request a la entidad User
	user := entities.User{
		IDPersona:     req.IDPersona,
		NombreUsuario: req.NombreUsuario,
		Password:      req.Password,
	}
	createdUser, err := c.UseCase.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdUser)
}
