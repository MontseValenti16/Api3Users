package routes

import (
	"API3/src/usuarios/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes configura los endpoints para login y registro de usuario.
func RegisterAuthRoutes(r *gin.Engine, ctrl *controllers.AuthController) {
	// Ruta para login
	authGroup := r.Group("/login")
	{
		authGroup.POST("/", ctrl.Login) // POST: Login y generación de token
	}
	// Ruta para registrar un usuario
	userGroup := r.Group("/usuario")
	{
		userGroup.POST("/registro", ctrl.Register) // POST: Registro de usuario
	}
}
