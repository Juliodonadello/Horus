package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorEventController struct{}

func (e SensorEventController) Save(c *gin.Context) {
	var SensorEventNuevo models.SensorEvent
	if err := c.ShouldBindJSON(&SensorEventNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := SensorEventNuevo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := SensorEventNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
