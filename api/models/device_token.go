package models

import (
	"api/conf_db"
	"crypto/rand"
	"encoding/hex"
	"errors"
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

// Create gets a device name and fills struct fields with a random token, the time and the time that the token expires
func (e *DeviceToken) Create(device string) error {
	e.Token = generateSecureToken(tokenLength)
	e.Device = device
	actualTime := time.Now()
	e.CreatedAt = actualTime
	e.ExpiresAt = actualTime.AddDate(0, 0, tokenDurationDays)
	return e.Validate()
}

// Write saves the token to the database and an in memory cache maps.
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

// FindByToken requires a token string, if exists returns a pointer to the token struct.
func FindByToken(token string) *DeviceToken {
	//client := conf_db.GetConfDB()
	//stmt := `SELECT device,token,created_at,expires_at FROM device_tokens WHERE token = $1;`
	//var oldToken DeviceToken
	//err := (*client).QueryRowx(stmt, token).StructScan(&oldToken)
	//if err != nil {
	//	return nil, err
	//}
	//delete(TokenCache, oldToken.Token)
	//delete(DeviceCache, oldToken.Device)
	return TokenCache[token]
}

// FindByDevice as device names are unique, gets a string with a device name and tries to find it at token cache.
func FindByDevice(device string) *DeviceToken {
	//client := conf_db.GetConfDB()
	//stmt := `SELECT device,token,created_at,expires_at FROM device_tokens WHERE device = $1;`
	//var oldToken DeviceToken
	//err := (*client).QueryRowx(stmt, device).StructScan(&oldToken)
	//if err != nil {
	//	return nil, err
	//}
	//delete(TokenCache, oldToken.Token)
	//delete(DeviceCache, oldToken.Device)
	return DeviceCache[device]
}

// Delete removes token from database and cache.
func (e *DeviceToken) Delete() error {
	client := conf_db.GetConfDB()
	stmt := `DELETE FROM device_tokens WHERE token = $1;`
	_, err := (*client).Exec(stmt, e.Token)
	if err != nil {
		return err
	}
	delete(TokenCache, e.Token)
	delete(DeviceCache, e.Device)
	return nil
}

// Validate checks token fields correctness
func (e DeviceToken) Validate() error {
	if strings.TrimSpace(e.Device) == "" || len(e.Device) > 50 {
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

// CheckToken auxiliar function takes a token string and tries to find it in cache.
func CheckToken(token string) (bool, string) {
	result := TokenCache[token]
	if result == nil {
		return false, ""
	}
	return true, result.Device
}

// generateSecureToken fills a length []byte array with random bits, and returns a encoded hex string.
func generateSecureToken(length int) string {
	if length >= 0 && length <= 32 {
		b := make([]byte, length)
		if _, err := rand.Read(b); err != nil {
			return ""
		}
		return hex.EncodeToString(b)
	}
	return ""
}

// RecoverTokens auxiliar function fills cache with tokens saved in database.
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
			//fmt.Printf("Se elimino token inválido %s\n", err.Error())
			continue
		}
		TokenCache[token.Token] = &token
		DeviceCache[token.Device] = &token
	}
	return cont, nil
}
