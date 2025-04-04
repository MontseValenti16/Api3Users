package handlers

import (
	"API3/src/mac_address/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompareMacsController struct {
	UseCase *application.CompareMacsUseCase
}

func NewCompareMacsController(useCase *application.CompareMacsUseCase) *CompareMacsController {
	return &CompareMacsController{UseCase: useCase}
}

func (c *CompareMacsController) CompareMacsHandler(ctx *gin.Context) {
	matchedDevices, err := c.UseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error comparando MACs"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"coinciden": matchedDevices})
}
