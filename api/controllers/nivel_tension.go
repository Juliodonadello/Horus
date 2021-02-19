package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TensionGeneradorController struct{}

func (e TensionGeneradorController) Save(c *gin.Context) {
	var tensionGeneradorNuevo models.TensionGenerador
	if err := c.ShouldBindJSON(&tensionGeneradorNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tensionGeneradorNuevo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := tensionGeneradorNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
