package models

import (
	"api/metrics_db"
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

var (
	InvalidGeneratorAlimentacion = errors.New("Alimentacion Model: Invalid Generator Name")
	InvalidTensionAlimentacion   = errors.New("Alimentacion Model: Invalid tension level value")
	InvalidCorrienteAlimentacion = errors.New("Alimentacion Model: Invalid corriente level value")
)

type Alimentacion struct {
	Generador string  `json:"generador" binding:"required"`
	Facilidad string  `json:"facilidad"`
	Tension   float32 `json:"tension" binding:"numeric"`
	Corriente float32 `json:"corriente" binding:"numeric"`
}

func (e Alimentacion) Write() (*Alimentacion, error) {
	client := metrics_db.GetTSDB()
	writeAPI := (*client).WriteAPIBlocking("ccic", "alimentacion-electrica")
	p := influxdb2.NewPoint("alimentacion",
		map[string]string{"generador": e.Generador, "facilidad": e.Facilidad},
		map[string]interface{}{"tension": e.Tension, "corriente": e.Corriente},
		time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Alimentacion) Validate() error {
	if len(e.Generador) <= 0 {
		return InvalidGeneratorAlimentacion
	}
	if e.Tension < 0.0 || e.Tension > 500.0 {
		return InvalidTensionAlimentacion
	}
	if e.Corriente < 0.0 || e.Corriente > 50.0 {
		return InvalidCorrienteAlimentacion
	}
	return nil
}
