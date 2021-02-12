package models

import (
	"api-horus/api/db"
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type NivelCombustible struct {
	Generador string  `json:"generador"`
	Nivel     float32 `json:"nivel"`
	Time      int64   `json:"timestamp"`
}

func (e NivelCombustible) Write() (*NivelCombustible, error) {
	client := db.GetTSDB()
	writeAPI := (*client).WriteAPIBlocking("ccic", "combustible-generador")
	p := influxdb2.NewPoint("litros",
		map[string]string{"generador": e.Generador},
		map[string]interface{}{"nivel": e.Nivel},
		time.Unix(e.Time, 0))
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
