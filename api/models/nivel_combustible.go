package models

import (
	"api/metrics_db"
	"context"
	"errors"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	ErrFooInvalidGeneradorNivelCombustible = errors.New("NivelCombustible Model: Invalid Generator Name")
	ErrFooInvalidNivelNivelCombustible     = errors.New("NivelCombustible Model: Invalid Gas level value")
	ErrFooInvalidTimeNivelCombustible      = errors.New("NivelCombustible Model: Invalid Time Value, out of +- 600 sec range")
	ErrFooInvalidCapacidadNivelCombustible = errors.New("NivelCombustible Model: Invalid Capacidad")
)

type NivelCombustible struct {
	Generador string  `json:"generador" binding:"required"`
	Nivel     float32 `json:"nivel" binding:"numeric"`
	Capacidad float32 `json:"capacidad" binding:"required"`
	Time      int64   `json:"timestamp"`
}

func (e NivelCombustible) Write() (*NivelCombustible, error) {
	client := metrics_db.GetTSDB()
	writeAPI := (*client).WriteAPIBlocking("ccic", "combustible-generador")
	p := influxdb2.NewPoint("litros",
		map[string]string{"generador": e.Generador},
		map[string]interface{}{
			"nivel":     e.Nivel,
			"capacidad": e.Capacidad},
		time.Unix(e.Time, 0))
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e NivelCombustible) Validate() error {
	if strings.TrimSpace(e.Generador) == "" {
		return ErrFooInvalidGeneradorNivelCombustible
	}
	if e.Capacidad < 0.0 || e.Capacidad < e.Nivel {
		return ErrFooInvalidCapacidadNivelCombustible
	}
	if e.Nivel < 0.0 {
		return ErrFooInvalidNivelNivelCombustible
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return ErrFooInvalidTimeNivelCombustible
	}
	return nil
}
