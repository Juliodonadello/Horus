package models

import (
	"api/event_db"
	"errors"
	"strings"
	"time"
)

var (
	ErrFooInvalidFacilityGPSEvent  = errors.New("SensorEventGPS Model: Invalid Facility for GPS Event")
	ErrFooInvalidLatitudeGPSEvent  = errors.New("SensorEventGPS Model: Invalid Latitude for Sensor Event")
	ErrFooInvalidLongitudeGPSEvent = errors.New("SensorEventGPS Model: Invalid Longitude for Sensor Event")
	ErrFooInvalidAltitudeGPSEvent  = errors.New("SensorEventGPS Model: Invalid Altitude for Sensor Event")
	ErrFooInvalidTimeGPSEvent      = errors.New("SensorEventGPS Model: Invalid Time for Sensor Event")
)

type SensorEvent_gps struct {
	Facilidad string  `json:"facilidad" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"numeric"`
	Longitude float32 `json:"longitude" binding:"numeric"`
	Altitude  float32 `json:"altitude" binding:"numeric"`
	Time      int64   `json:"timestamp"`
}

func (e SensorEvent_gps) Write() error {
	client := event_db.GetEventDB()
	stmt := `INSERT INTO gps_events (facilidad, latitude, longitude, altitude, gfh)
				VALUES ($1, $2, $3, $4, $5);`
	_, err := (*client).Exec(stmt, e.Facilidad, e.Latitude, e.Longitude, e.Altitude, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (e SensorEvent_gps) Validate() error {
	if strings.TrimSpace(e.Facilidad) == "" {
		return ErrFooInvalidFacilityGPSEvent
	}
	if e.Latitude < -90 || e.Latitude > 90 {
		return ErrFooInvalidLatitudeGPSEvent
	}
	if e.Longitude < -180 || e.Longitude > 180 {
		return ErrFooInvalidLongitudeGPSEvent
	}
	if e.Altitude < -1000 || e.Altitude > 10000 {
		return ErrFooInvalidAltitudeGPSEvent
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return ErrFooInvalidTimeGPSEvent
	}
	return nil
}
