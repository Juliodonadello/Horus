package models

import (
	"api-horus/api/db"
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

var (
	InvalidFacilityNivelCombustible = errors.New("LongitudCola Model: Invalid Generator Name")
	InvalidNivelNivelCombustible    = errors.New("LongitudCola Model: Invalid Gas level value")
	InvalidTimeNivelCombustible     = errors.New("LongitudCola Model: Invalid Time Value, out of +- 600 sec range")
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

func (e NivelCombustible) Validate() error {
	if len(e.Generador) <= 0 {
		return InvalidFacilityNivelCombustible
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
