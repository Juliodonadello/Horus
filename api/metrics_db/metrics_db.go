package metrics_db

import (
	"context"
	"errors"
	"github.com/influxdata/influxdb-client-go/v2"
	"os"
)

var host, token = os.Getenv("METRIC_DB_HOST"), os.Getenv("METRIC_DB_TOKEN")

var connectionError = errors.New("No se puede conectar a la TSDB InfluxDB")
var tsdb *influxdb2.Client

func Init() error {
	client := influxdb2.NewClient(host, token)
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
