package server

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)
import "api/controllers"

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(secure.New(secure.Config{
		AllowedHosts:          []string{"localhost"},
		SSLRedirect:           false,
		SSLHost:               "",
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", Pong)
	v1 := router.Group("v1")
	{
		metricsGroup := v1.Group("metricas")
		{
			estadoServicio := new(controllers.EstadoServicioController)
			longitudCola := new(controllers.LongitudColaController)
			nivelCombustible := new(controllers.NivelCombustibleController)
			tensionGenerador := new(controllers.TensionGeneradorController)

			metricsGroup.POST("estado-servicio", estadoServicio.Save)
			metricsGroup.POST("longitud-cola", longitudCola.Save)
			metricsGroup.POST("nivel-combus", nivelCombustible.Save)
			metricsGroup.POST("nivel-tension", tensionGenerador.Save)

		}
		eventsGroup := v1.Group("eventos")
		{
			mensMilEvent := new(controllers.MensMilEventController)
			eventsGroup.POST("mens-mil", mensMilEvent.Save)
			eventsGroup.DELETE("mens-mil", mensMilEvent.ClearRegisters)
		}
	}
	return router
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
