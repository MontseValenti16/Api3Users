package handlers

import (
	"API3/src/mac_address/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScanNetworkController struct {
	UseCase *application.ScanNetworkUseCase
}

func NewScanNetworkController(useCase *application.ScanNetworkUseCase) *ScanNetworkController {
	return &ScanNetworkController{UseCase: useCase}
}

func (c *ScanNetworkController) ScanNetworkHandler(ctx *gin.Context) {
	devices, err := c.UseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error escaneando la red"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"devices": devices})
}
