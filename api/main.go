package main

import (
	"api-horus/api/db"
	"api-horus/api/server"
)

// https://gin-gonic.com/docs/examples/custom-http-config/
// https://github.com/vsouza/go-gin-boilerplate
func main() {
	s := server.Init()
	err := db.Init()
	if err != nil {
		panic("No se pudo conectar con Metricas-DB")
	}
	_ = s.ListenAndServe() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
