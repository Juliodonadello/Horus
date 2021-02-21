package models

import (
	"api/conf_db"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	tokenLengthEnv       = os.Getenv("TOKEN_LENGTH")
	tokenDurationDaysEnv = os.Getenv("TOKEN_DURATION_DAYS")
	tokenDurationDays, _ = strconv.Atoi(tokenDurationDaysEnv)
	tokenLength, _       = strconv.Atoi(tokenLengthEnv)
	TokenCache           = make(map[string]*DeviceToken, 20)
	DeviceCache          = make(map[string]*DeviceToken, 20)
	invalidDevice        = errors.New("DeviceToken Model: Invalid Device Name")
	invalidToken         = errors.New("DeviceToken Model: Invalid Token")
	invalidCreationTime  = errors.New("DeviceToken Model: Invalid created at")
	invalidDurationDays  = errors.New("DeviceToken Model: Invalid duration days")
	deviceExists         = errors.New("DeviceToken Model: Device exists")
)

type DeviceToken struct {
	Device    string    `json:"device" db:"device"`
	Token     string    `json:"token" db:"token"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

func (e *DeviceToken) Create(device string) {
	e.Token = generateSecureToken(tokenLength)
	e.Device = device
	actualTime := time.Now()
	e.CreatedAt = actualTime
	e.ExpiresAt = actualTime.AddDate(0, 0, tokenDurationDays)
}

func (e DeviceToken) Write() error {
	client := conf_db.GetConfDB()
	stmt := `INSERT INTO device_tokens (token, device, created_at,expires_at) VALUES ($1, $2, $3, $4);`
	_, err := (*client).Exec(stmt, e.Token, e.Device, e.CreatedAt, e.ExpiresAt)
	if err != nil {
		return err
	}
	TokenCache[e.Token] = &e
	DeviceCache[e.Device] = &e
	return nil
}

func (e DeviceToken) Validate() error {
	if strings.TrimSpace(e.Device) == "" || len(e.Device) > 50 {
		fmt.Println(e.Device)
		return invalidDevice
	}
	if strings.TrimSpace(e.Token) == "" {
		return invalidToken
	}
	today := time.Now()
	if e.CreatedAt.After(today) {
		return invalidCreationTime
	}
	if int(e.ExpiresAt.Sub(e.CreatedAt).Hours()/24) != tokenDurationDays {
		return invalidDurationDays
	}
	if DeviceCache[e.Device] != nil {
		return deviceExists
	}
	return nil
}

func CheckToken(token string) (bool, string) {
	result := TokenCache[token]
	if result == nil {
		return false, ""
	}
	return true, result.Device
}

func generateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func RecoverTokens() (int, error) {
	client := conf_db.GetConfDB()
	var cont int
	stmt := `SELECT device,token,created_at,expires_at FROM device_tokens;`
	rows, err := (*client).Queryx(stmt)
	if err != nil {
		return cont, err
	}
	for rows.Next() {
		var token DeviceToken
		err := rows.StructScan(&token)
		cont++
		if err != nil {
			return cont, err
		}
		err = token.Validate()
		if err != nil {
			cont--
			fmt.Println("Se eliminó token inválido")
			continue
		}
		TokenCache[token.Token] = &token
		DeviceCache[token.Device] = &token
	}
	return cont, nil
}
