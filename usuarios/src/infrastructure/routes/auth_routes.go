package routes

import (
	"API3/usuarios/src/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes configura el endpoint para el login y la generación del token JWT.
func RegisterAuthRoutes(r *gin.Engine, ctrl *controllers.AuthController) {
	authGroup := r.Group("/login")
	{
		authGroup.POST("/", ctrl.Login) // POST: Login y generación de token
	}
}
