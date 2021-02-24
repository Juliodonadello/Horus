package controllers

import (
	"api/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	grafanaHost         = os.Getenv("GRAFANA_HOST")
	grafanaPort         = os.Getenv("GRAFANA_PORT")
	externalAuthProblem = errors.New("DeviceTokenController: Can't reach dashboard user API")
)

type DeviceTokenController struct{}

type LoginData struct {
	Username string `form:"user" binding:"required"`
	Password string `form:"pass" binding:"required"`
	Device   string `form:"device"`
}

func (d DeviceTokenController) GetToken(c *gin.Context) {
	var loginData LoginData
	err := c.BindQuery(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	isAdmin, err := userIsAdmin(loginData)
	var newToken models.DeviceToken
	if isAdmin {
		err = newToken.Create(loginData.Device)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
			return
		}
		err = newToken.Write()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"device": newToken.Device, "token": newToken.Token})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": externalAuthProblem.Error(), "status": http.StatusInternalServerError})
	return
}

func (d DeviceTokenController) RevokeToken(c *gin.Context) {
	var loginData LoginData
	err := c.BindQuery(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	isAdmin, err := userIsAdmin(loginData)
	if isAdmin {
		token := c.Query("token")
		device := c.Query("device")
		var oldToken *models.DeviceToken
		var err error
		if token == "" && device == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "neither token nor device", "status": http.StatusBadRequest})
			return
		}
		if token == "" {
			oldToken = models.FindByDevice(device)
			if oldToken != nil {
				err = oldToken.Delete()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
					return
				}
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "dispositivo no hallado", "status": http.StatusNotFound})
				return
			}
		} else {
			oldToken = models.FindByToken(token)
			if oldToken != nil {
				err = oldToken.Delete()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
					return
				}
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "token no hallado", "status": http.StatusNotFound})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"device": oldToken.Device, "token": oldToken.Token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
	}
}

func (d DeviceTokenController) ListTokens(c *gin.Context) {
	var loginData LoginData
	err := c.BindQuery(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	isAdmin, err := userIsAdmin(loginData)
	if isAdmin {
		authorizedDevices := make(map[string]string)
		for k, v := range models.DeviceCache {
			authorizedDevices[k] = v.Token
		}
		jsonString, err := json.Marshal(authorizedDevices)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
			return
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		_, err = c.Writer.Write(jsonString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
			return
		}
	}
}

func userIsAdmin(loginData LoginData) (bool, error) {
	url := fmt.Sprintf("https://%s:%s@%s:%s/api/user", loginData.Username, loginData.Password, grafanaHost, grafanaPort)
	resp, err := http.Get(url)
	if err != nil {
		//c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
		return false, err
	}
	var objmap map[string]json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		//c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error(), "status": http.StatusServiceUnavailable})
		return false, err
	}
	var isAdmin bool
	if objmap["isGrafanaAdmin"] == nil {
		//c.JSON(http.StatusOK, gin.H{"message": "invalid username or password"})
		return false, errors.New("invalid username or password")
	}
	err = json.Unmarshal(objmap["isGrafanaAdmin"], &isAdmin)
	if err != nil {
		//c.JSON(http.StatusServiceUnavailable, gin.H{"message": err.Error(), "status": http.StatusServiceUnavailable})
		return false, err
	}
	return isAdmin, nil
}
