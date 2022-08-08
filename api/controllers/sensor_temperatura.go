package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemperatureEventController struct{}

func (e TemperatureEventController) Save(c *gin.Context) {
	var TemperatureEventNuevo models.TemperatureEvent
	if err := c.ShouldBindJSON(&TemperatureEventNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := TemperatureEventNuevo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := TemperatureEventNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
