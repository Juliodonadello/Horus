package models

import (
	"api/event_db"
	"errors"
	"strings"
	"time"
)

var (
	ErrFooInvalidFacilitySensorEvent = errors.New("SensorEvent Model: Invalid Facility for Sensor Event")
	ErrFooInvalidEventoSensorEvent   = errors.New("SensorEvent Model: Invalid Event for Sensor Event")
)

type SensorEvent struct {
	Facilidad string `json:"facilidad" binding:"required"`
	Evento    string `json:"evento" binding:"required"`
	Valor     *bool  `json:"valor" binding:"required"`
}

func (e SensorEvent) Write() error {
	client := event_db.GetEventDB()
	stmt := `INSERT INTO sensor_events (facilidad, evento, valor, gfh) 
				VALUES ($1, $2, $3, $4) 
				ON CONFLICT (facilidad, evento) 
				DO UPDATE SET
					valor = EXCLUDED.valor,
					gfh = EXCLUDED.gfh;`
	_, err := (*client).Exec(stmt, e.Facilidad, e.Evento, e.Valor, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (e SensorEvent) Validate() error {
	if strings.TrimSpace(e.Facilidad) == "" {
		return ErrFooInvalidFacilitySensorEvent
	}
	if strings.TrimSpace(e.Evento) == "" {
		return ErrFooInvalidEventoSensorEvent
	}
	return nil
}
