package infraestructure

import (
	"API3/src/mac_address/application"
	"API3/src/mac_address/infraestructure/handlers"
	infraestructure "API3/src/mac_address/infraestructure/repository"
	"API3/src/mac_address/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	repo := infraestructure.NewNetworkRepository()
	compare := application.NewCompareMacsUseCase(repo)
	scan := application.NewScanNetworkUseCase(repo)
	handlerCompare := handlers.NewCompareMacsController(compare)
	handleScan := handlers.NewScanNetworkController(scan)
	routes.SetupRouter(r, handleScan, handlerCompare)

}
