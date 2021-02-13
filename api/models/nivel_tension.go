package models

import (
	"api/tsdb"
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

var (
	InvalidGeneratorTensionGenerador = errors.New("TensionGenerador Model: Invalid Generator Name")
	InvalidNivelTensionGenerador     = errors.New("TensionGenerador Model: Invalid tension level value")
	InvalidTimeTensionGenerador      = errors.New("TensionGenerador Model: Invalid Time Value, out of +- 600 sec range")
)

type TensionGenerador struct {
	Generador string  `json:"generador" binding:"required"`
	Tension   float32 `json:"tension" binding:"required"`
	Time      int64   `json:"timestamp"`
}

func (e TensionGenerador) Write() (*TensionGenerador, error) {
	client := tsdb.GetTSDB()
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

func (e TensionGenerador) Validate() error {
	if len(e.Generador) <= 0 {
		return InvalidGeneratorTensionGenerador
	}
	if e.Tension < 0.0 || e.Tension > 500.0 {
		return InvalidNivelTensionGenerador
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return InvalidTimeTensionGenerador
	}
	return nil
}
