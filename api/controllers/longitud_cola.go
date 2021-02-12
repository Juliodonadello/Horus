package controllers

import (
	"api-horus/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LongitudColaController struct{}

func (e LongitudColaController) Save(c *gin.Context) {
	var longitudColaNueva models.LongitudCola
	if err := c.ShouldBindJSON(&longitudColaNueva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := longitudColaNueva.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := longitudColaNueva.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
