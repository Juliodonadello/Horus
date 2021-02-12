package models

import (
	"api-horus/api/db"
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type EstadoServicio struct {
	Facilidad string `json:"facilidad" binding:"required"`
	Estado    string `json:"estado" binding:"required"`
	Time      int64  `json:"timestamp" binding:"required"`
}

func (e EstadoServicio) Write() (*EstadoServicio, error) {
	client := db.GetTSDB()
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
