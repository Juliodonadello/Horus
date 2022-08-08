package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorGPSController struct{}

func (e SensorGPSController) Save(c *gin.Context) {
	var SensorGPSNuevo models.SensorEvent_gps
	if err := c.ShouldBindJSON(&SensorGPSNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := SensorGPSNuevo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := SensorGPSNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
