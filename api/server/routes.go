package server

import (
	"api/middleware"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"os"
)
import "api/controllers"

var (
	grafanaHost = os.Getenv("GRAFANA_HOST")
	grafanaPort = os.Getenv("GRAFANA_PORT")
)

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
	healthReport := new(controllers.HealthCheckController)
	router.GET("health-check", healthReport.HealthCheckReport)
	v1 := router.Group("v1")
	{

		metricsGroup := v1.Group("metricas")
		metricsGroup.Use(middleware.TokenDeviceAuth())
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
		eventsGroup.Use(middleware.TokenDeviceAuth())
		{
			mensMilEvent := new(controllers.MensMilEventController)
			eventsGroup.POST("mens-mil", mensMilEvent.Save)
			eventsGroup.DELETE("mens-mil", mensMilEvent.ClearRegisters)
		}
		tokenGroup := v1.Group("token")
		{
			deviceToken := new(controllers.DeviceTokenController)
			tokenGroup.GET("generate", deviceToken.GetToken)
			tokenGroup.DELETE("revoke", deviceToken.RevokeToken)
			tokenGroup.GET("list", deviceToken.ListTokens)
		}
	}
	return router
}
