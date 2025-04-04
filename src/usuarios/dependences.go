package usuarios

import (
	auth_usecase "API3/src/usuarios/application"
	"API3/src/usuarios/infrastructure/controllers"
	"API3/src/usuarios/infrastructure/repositories"
	"API3/src/usuarios/infrastructure/routes"

	"github.com/gin-gonic/gin"
)


func Init(r *gin.Engine) {
	db := repositories.NewUserRepository()
	createUseCase := auth_usecase.NewAuthUseCase(db)
	controller := controllers.NewAuthController(createUseCase)
	routes.RegisterAuthRoutes(r, controller)
}
