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
	EstadosValid          = []string{"en servicio", "servicio limitado", "fuera de servicio"}
	ErrFooInvalidFacility = errors.New("EstadoServicio Model: Invalid Facility Name")
	ErrFooInvalidState    = errors.New("EstadoServicio Model: Invalid State Name")
	ErrFooInvalidTime     = errors.New("EstadoServicio Model: Invalid Time Value, out of +- 600 sec range")
)

type EstadoServicio struct {
	Facilidad string `json:"facilidad" binding:"required"`
	Estado    string `json:"estado" binding:"required"`
	Time      int64  `json:"timestamp"`
}

func (e EstadoServicio) Write() (*EstadoServicio, error) {
	client := metrics_db.GetTSDB()
	writeAPI := (*client).WriteAPIBlocking("ccic", "estado-servicio")
	p := influxdb2.NewPoint("estado",
		map[string]string{"facilidad": e.Facilidad},
		map[string]interface{}{"estado": e.Estado},
		time.Unix(e.Time, 0))
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e EstadoServicio) Validate() error {
	if strings.TrimSpace(e.Facilidad) == "" {
		return ErrFooInvalidFacility
	}
	var estado string
	for i := 0; i < len(EstadosValid); i++ {
		if EstadosValid[i] == e.Estado {
			estado = e.Estado
			break
		}
	}
	if len(estado) == 0 {
		return ErrFooInvalidState
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return ErrFooInvalidTime
	}
	return nil
}
