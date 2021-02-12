package controllers

import (
	"api-horus/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EstadoServicioController struct{}

func (e EstadoServicioController) Save(c *gin.Context) {
	var estadoServicioNuevo models.EstadoServicio
	if err := c.ShouldBindJSON(&estadoServicioNuevo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	// Aca validar el contenido de la struct
	_, err := estadoServicioNuevo.Write()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
