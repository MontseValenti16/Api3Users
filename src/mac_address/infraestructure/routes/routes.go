package routes

import (
	"API3/src/mac_address/infraestructure/handlers"

	"github.com/gin-gonic/gin"
)


func SetupRouter(r *gin.Engine, controllerScan *handlers.ScanNetworkController,  controllerCompare *handlers.CompareMacsController){
	r.GET("/scan", controllerScan.ScanNetworkHandler)
	r.GET("/compare", controllerCompare.CompareMacsHandler )
}
