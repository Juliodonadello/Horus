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
	EstadosValid    = []string{"en servicio", "servicio limitado", "fuera de servicio"}
	InvalidFacility = errors.New("EstadoServicio Model: Invalid Facility Name")
	InvalidState    = errors.New("EstadoServicio Model: Invalid State Name")
	InvalidTime     = errors.New("EstadoServicio Model: Invalid Time Value, out of +- 600 sec range")
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
		return InvalidFacility
	}
	var estado string
	for i := 0; i < len(EstadosValid); i++ {
		if EstadosValid[i] == e.Estado {
			estado = e.Estado
			break
		}
	}
	if len(estado) == 0 {
		return InvalidState
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return InvalidTime
	}
	return nil
}
