package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MensMilEventController struct{}

func (e MensMilEventController) Save(c *gin.Context) {
	var MensMilEventNuevo models.MensMilEvent
	if err := c.ShouldBindJSON(&MensMilEventNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MensMilEventNuevo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := MensMilEventNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (e MensMilEventController) ClearRegisters(c *gin.Context) {
	var MensMilEventNuevo models.MensMilEvent
	_, err := MensMilEventNuevo.Clear()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
