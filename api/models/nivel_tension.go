package models

import (
	"api-horus/api/db"
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type TensionGenerador struct {
	Generador string  `json:"generador"`
	Tension   float32 `json:"tension"`
	Time      int64   `json:"timestamp"`
}

func (e TensionGenerador) Write() (*TensionGenerador, error) {
	client := db.GetTSDB()
	writeAPI := (*client).WriteAPIBlocking("ccic", "tension-generador")
	p := influxdb2.NewPoint("volts",
		map[string]string{"generador": e.Generador},
		map[string]interface{}{"tension": e.Tension},
		time.Unix(e.Time, 0))
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
