package metrics_db

import (
	"context"
	"errors"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var host, token = os.Getenv("METRIC_DB_HOST"), os.Getenv("METRIC_DB_TOKEN")

var ErrFooConnectionError = errors.New("no se puede conectar a la TSDB InfluxDB")
var tsdb *influxdb2.Client

func Init() error {
	client := influxdb2.NewClient(host, token)
	server_status, err := client.Health(context.Background())
	if err != nil {
		return err
	}
	if server_status.Status == "fail" {
		return ErrFooConnectionError
	}
	tsdb = &client
	return nil
}

func GetTSDB() *influxdb2.Client {
	return tsdb
}
