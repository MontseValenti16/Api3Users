package registro

import (
	"API3/registro/src/application"
	"API3/registro/src/infrastructure/controllers"
	"API3/registro/src/infrastructure/repositories"
	"API3/registro/src/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	repo := repositories.NewUserRepository()
	uc := user_usecase.NewUserUseCase(repo)
	ctrl := controllers.NewUserController(uc)
	routes.RegisterUserRoutes(r, ctrl)
}
