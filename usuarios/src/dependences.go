package usuarios

import (
	"API3/usuarios/src/application"
	"API3/usuarios/src/infrastructure/controllers"
	"API3/usuarios/src/infrastructure/repositories"
	"API3/usuarios/src/infrastructure/routes"

	"github.com/gin-gonic/gin"
)


func Init(r *gin.Engine) {
	repo := repositories.NewUserRepository()
	uc := auth_usecase.NewAuthUseCase(repo)
	ctrl := controllers.NewAuthController(uc)
	routes.RegisterAuthRoutes(r, ctrl)
}
