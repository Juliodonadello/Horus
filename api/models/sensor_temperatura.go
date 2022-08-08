package models

import (
	"api/event_db"
	"errors"
	"strings"
	"time"
)

var (
	ErrFooInvalidFacilityTempEvent    = errors.New("TemperatureEvent Model: Invalid Facility for Temperature Event")
	ErrFooInvalidTemperatureTempEvent = errors.New("TemperatureEvent Model: Invalid Temperature for Temperature Event")
	ErrFooInvalidTimeTempEvent        = errors.New("TemperatureEvent Model: Invalid Time for Temperature Event")
)

type TemperatureEvent struct {
	Facilidad   string  `json:"facilidad" binding:"required"`
	Temperature float32 `json:"temperature" binding:"numeric"`
	Time        int64   `json:"timestamp"`
}

func (e TemperatureEvent) Write() error {
	client := event_db.GetEventDB()
	stmt := `INSERT INTO temperature_events (facilidad, temperature, gfh)
				VALUES ($1, $2, $3);`
	_, err := (*client).Exec(stmt, e.Facilidad, e.Temperature, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (e TemperatureEvent) Validate() error {
	if strings.TrimSpace(e.Facilidad) == "" {
		return ErrFooInvalidFacilityTempEvent
	}
	if e.Temperature < -50 || e.Temperature > 150 {
		return ErrFooInvalidTemperatureTempEvent
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return ErrFooInvalidTimeTempEvent
	}
	return nil
}
