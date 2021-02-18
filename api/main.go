package main

import (
	"api/event_db"
	"api/metrics_db"
	"api/server"
	"time"
)

// https://gin-gonic.com/docs/examples/custom-http-config/
// https://github.com/vsouza/go-gin-boilerplate
func main() {
	time.Sleep(10 * time.Second)
	s := server.Init()
	err := metrics_db.Init()
	if err != nil {
		panic("No se pudo conectar con Metricas-DB")
	}
	err = event_db.Init()
	if err != nil {
		panic("No se pudo conectar con Event-DB")
	}
	_ = s.ListenAndServe() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
