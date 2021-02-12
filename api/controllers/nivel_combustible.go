package controllers

import (
	"api-horus/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NivelCombustibleController struct{}

func (e NivelCombustibleController) Save(c *gin.Context) {
	var nivelCombustibleNuevo models.NivelCombustible
	if err := c.ShouldBindJSON(&nivelCombustibleNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := nivelCombustibleNuevo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := nivelCombustibleNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
