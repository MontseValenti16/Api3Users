package main

import (
	"API3/core/middleware"
	"API3/core/mysql"
	registro "API3/registro/src"
	usuarios "API3/usuarios/src"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql.InitDB()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	registro.Init(router)
	usuarios.Init(router)

	log.Println("Servidor corriendo en el puerto 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
