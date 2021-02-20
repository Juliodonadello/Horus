package controllers

import (
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
}

func (d DeviceTokenController) GetToken(c *gin.Context) {
	var loginData LoginData
	err := c.BindQuery(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
	}
	url := fmt.Sprintf("http://%s:%s@%s:%s/api/user", loginData.Username, loginData.Password, grafanaHost, grafanaPort)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
	}
	sb := string(body)
	fmt.Println(sb)
}
