package routes

import (
	"API3/registro/src/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes configura los endpoints para el módulo de registro.
func RegisterUserRoutes(r *gin.Engine, ctrl *controllers.UserController) {
	userGroup := r.Group("/registro")
	{
		userGroup.POST("/", ctrl.RegisterUser)       // POST: Registrar usuario
		userGroup.GET("/:id", ctrl.GetUser)            // GET: Obtener usuario por id
		userGroup.GET("/", ctrl.GetUsers)              // GET: Listar todos los usuarios
		userGroup.PUT("/", ctrl.UpdateUser)            // PUT: Actualizar usuario
		userGroup.DELETE("/:id", ctrl.DeleteUser)       // DELETE: Eliminar usuario
	}
}
