package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	grafanaHost = os.Getenv("GRAFANA_HOST")
	grafanaPort = os.Getenv("GRAFANA_PORT")
)

type DeviceTokenController struct{}

type LoginData struct {
	Username string `form:"user"`
	Password string `form:"pass"`
	Device   string `form:"device"`
}

func (d DeviceTokenController) GetToken(c *gin.Context) {
	var loginData LoginData
	err := c.BindQuery(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	url := fmt.Sprintf("http://%s:%s@%s:%s/api/user", loginData.Username, loginData.Password, grafanaHost, grafanaPort)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
		return
	}
	var objmap map[string]json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
		return
	}
	var isAdmin bool
	if objmap["isGrafanaAdmin"] == nil {
		c.JSON(http.StatusOK, gin.H{"message": "invalid username or password"})
		return
	}
	err = json.Unmarshal(objmap["isGrafanaAdmin"], &isAdmin)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": err.Error(), "status": http.StatusServiceUnavailable})
		return
	}
	var newToken models.DeviceToken
	if isAdmin {
		newToken.Create(loginData.Device)
		err = newToken.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
			return
		}
		err = newToken.Write()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"device": newToken.Device, "token": newToken.Token})
	return
}
