package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EstadoServicioController struct{}

// Save guarda un update del estado de servicio de la facilidad
func (e EstadoServicioController) Save(c *gin.Context) {
	var estadoServicioNuevo models.EstadoServicio
	if err := c.ShouldBindJSON(&estadoServicioNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	if err := estadoServicioNuevo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := estadoServicioNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
