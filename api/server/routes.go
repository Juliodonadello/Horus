package server

import (
	"api/controllers"
	"api/middleware"
	"net/http"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
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
	router.LoadHTMLFiles("./docs/index.html")
	healthReport := new(controllers.HealthCheckController)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("health-check", healthReport.HealthCheckReport)
	v1 := router.Group("v1")
	{
		metricsGroup := v1.Group("metricas")
		metricsGroup.Use(middleware.TokenDeviceAuth())
		{
			estadoServicio := new(controllers.EstadoServicioController)
			longitudCola := new(controllers.LongitudColaController)
			nivelCombustible := new(controllers.NivelCombustibleController)
			tensionGenerador := new(controllers.AlimentacionController)

			metricsGroup.POST("estado-servicio", estadoServicio.Save)
			metricsGroup.POST("longitud-cola", longitudCola.Save)
			metricsGroup.POST("nivel-combus", nivelCombustible.Save)
			metricsGroup.POST("alimentacion", tensionGenerador.Save)

		}
		eventsGroup := v1.Group("eventos")
		eventsGroup.Use(middleware.TokenDeviceAuth())
		{
			mensMilEvent := new(controllers.MensMilEventController)
			sensorEvent := new(controllers.SensorEventController)
			sensorGPSEvent := new(controllers.SensorGPSController)

			eventsGroup.POST("mens-mil", mensMilEvent.Save)
			eventsGroup.DELETE("mens-mil", mensMilEvent.ClearRegisters)
			eventsGroup.POST("sensor_bool", sensorEvent.Save)
			eventsGroup.POST("geolocalizacion", sensorGPSEvent.Save)
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
