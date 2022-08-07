package models

import (
	"api/event_db"
	"errors"
	"strings"
	"time"
)

var (
	InvalidFacilitySensorEvent = errors.New("SensorEvent Model: Invalid Facility for Sensor Event")
	InvalidEventoSensorEvent   = errors.New("SensorEvent Model: Invalid Event for Sensor Event")
)

type SensorEvent_gps struct {
	Facilidad string `json:"facilidad" binding:"required"`
	Evento    string `json:"evento" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
	Altitude  float32 `json:"altitude" binding:"numeric"`
	Time      int64   `json:"timestamp"`
    }

func (e SensorEvent_gps) Write() error {
	client := event_db.GetEventDB()
	stmt := `INSERT INTO sensor_events (facilidad, evento, latitude, longitude, altitude, time)
				VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := (*client).Exec(stmt, e.Facilidad, e.Evento, e.Latitude, e.Longitude, e.Altitude, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (e SensorEvent) Validate() error {
	if strings.TrimSpace(e.Facilidad) == "" {
		return InvalidFacilitySensorEvent
	}
	if strings.TrimSpace(e.Evento) == "" {
		return InvalidEventoSensorEvent
	}
	return nil
}
