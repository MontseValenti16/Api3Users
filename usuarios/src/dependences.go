package usuarios

import (
	"API3/usuarios/src/application"
	"API3/usuarios/src/infrastructure/controllers"
	"API3/usuarios/src/infrastructure/repositories"
	"API3/usuarios/src/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

// Init inicializa el módulo de usuarios (login): crea repositorio, caso de uso, controlador y registra la ruta.
func Init(r *gin.Engine) {
	repo := repositories.NewUserRepository()
	uc := auth_usecase.NewAuthUseCase(repo)
	ctrl := controllers.NewAuthController(uc)
	routes.RegisterAuthRoutes(r, ctrl)
}
