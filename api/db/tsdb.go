package db

import (
	"context"
	"errors"
	"github.com/influxdata/influxdb-client-go/v2"
)

var connectionError = errors.New("No se puede conectar a la TSDB InfluxDB")
var tsdb *influxdb2.Client

func Init() error {
	client := influxdb2.NewClient("http://localhost:8086", "u8C0TQ4SOAQ6za2BLZO-grxyIg5NT-WEQGNE9O6pYy73rPt9t0MNx8trbAgXEdyrwXYQ9dCRs5-zr3BMeeVufw==")
	server_status, err := client.Health(context.Background())
	if err != nil {
		return err
	}
	if server_status.Status == "fail" {
		return connectionError
	}
	tsdb = &client
	return nil
}

func GetTSDB() *influxdb2.Client {
	return tsdb
}
