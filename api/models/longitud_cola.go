package models

import (
	"api/metrics_db"
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"strings"
	"time"
)

var (
	InvalidFacilityLongitudCola  = errors.New("LongitudCola Model: Invalid Facility Name")
	InvalidQueueLongLongitudCola = errors.New("LongitudCola Model: Invalid Queue Long value")
	InvalidTimeLongitudCola      = errors.New("LongitudCola Model: Invalid Time Value, out of +- 600 sec range")
)

type LongitudCola struct {
	Facilidad string `json:"facilidad" binding:"required"`
	LongCola  *int   `json:"long_cola" binding:"required"`
	Time      int64  `json:"timestamp"`
}

func (e LongitudCola) Write() (*LongitudCola, error) {
	client := metrics_db.GetTSDB()
	writeAPI := (*client).WriteAPIBlocking("ccic", "colas-espera")
	p := influxdb2.NewPoint("long_cola",
		map[string]string{"facilidad": e.Facilidad},
		map[string]interface{}{"long_cola": (*e.LongCola)},
		time.Unix(e.Time, 0))
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e LongitudCola) Validate() error {
	if strings.TrimSpace(e.Facilidad) == "" {
		return InvalidFacilityLongitudCola
	}
	if (*e.LongCola) < 0 {
		return InvalidQueueLongLongitudCola
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return InvalidTimeLongitudCola
	}
	return nil
}
