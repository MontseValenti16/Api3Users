package infrastructure

import (
	user_usecase "API3/src/registro/application"
	"API3/src/registro/infrastructure/controllers"
	"API3/src/registro/infrastructure/repositories"
	"API3/src/registro/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	db := repositories.NewPersonaRepository()
	createUseCase := user_usecase.NewUserUseCase(db)
	controller := controllers.NewUserController(createUseCase)
	routes.RegisterUserRoutes(r, controller)
}
