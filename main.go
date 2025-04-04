package main

import (
	"API3/core/middleware"
	"API3/src/mac_address/infraestructure"
	"API3/src/registro/infrastructure"
	"API3/src/usuarios"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	infraestructure.Init(r)
	infrastructure.Init(r)
	usuarios.Init(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
