package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlimentacionController struct{}

func (e AlimentacionController) Save(c *gin.Context) {
	var alimentacionNueva models.Alimentacion
	if err := c.ShouldBindJSON(&alimentacionNueva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := alimentacionNueva.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := alimentacionNueva.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
