package models

import (
	"api-horus/api/db"
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type LongitudCola struct {
	Facilidad string `json:"facilidad"`
	LongCola  int    `json:"long_cola"`
	Time      int64  `json:"timestamp"`
}

func (e LongitudCola) Write() (*LongitudCola, error) {
	client := db.GetTSDB()
	writeAPI := (*client).WriteAPIBlocking("ccic", "colas-espera")
	p := influxdb2.NewPoint("long_cola",
		map[string]string{"facilidad": e.Facilidad},
		map[string]interface{}{"long_cola": e.LongCola},
		time.Unix(e.Time, 0))
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
