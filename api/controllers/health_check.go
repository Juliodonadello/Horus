package controllers

import (
	"api/event_db"
	"api/metrics_db"
	"api/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func (e HealthCheckController) HealthCheckReport(c *gin.Context) {
	tsdb := metrics_db.GetTSDB()
	tsdb_status, err := (*tsdb).Health(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), "db_status": http.StatusInternalServerError,
		})
		return
	}
	db := event_db.GetEventDB()
	db_status := "up"
	if err := db.Ping(); err != nil {
		db_status = "down"
	}
	url := fmt.Sprintf("https://%s:%s/api/health", grafanaHost, grafanaPort)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), "db_status": http.StatusInternalServerError,
		})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), "db_status": http.StatusInternalServerError,
		})
		return
	}
	var objmap map[string]json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), "db_status": http.StatusInternalServerError,
		})
		return
	}
	c.JSON(200, gin.H{
		"tsdb":    tsdb_status.Status,
		"db":      db_status,
		"grafana": string(objmap["database"]),
		"tokens":  len(models.TokenCache),
	})
}
