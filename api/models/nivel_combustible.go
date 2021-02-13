package models

import (
	"api/tsdb"
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"strings"
	"time"
)

var (
	InvalidGeneradorNivelCombustible = errors.New("NivelCombustible Model: Invalid Generator Name")
	InvalidNivelNivelCombustible     = errors.New("NivelCombustible Model: Invalid Gas level value")
	InvalidTimeNivelCombustible      = errors.New("NivelCombustible Model: Invalid Time Value, out of +- 600 sec range")
)

type NivelCombustible struct {
	Generador string  `json:"generador" binding:"required"`
	Nivel     float32 `json:"nivel" binding:"required"`
	Time      int64   `json:"timestamp"`
}

func (e NivelCombustible) Write() (*NivelCombustible, error) {
	client := tsdb.GetTSDB()
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

func (e NivelCombustible) Validate() error {
	if strings.TrimSpace(e.Generador) == "" {
		return InvalidGeneradorNivelCombustible
	}
	if e.Nivel < 0.0 {
		return InvalidNivelNivelCombustible
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return InvalidTimeNivelCombustible
	}
	return nil
}
