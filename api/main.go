package main

import (
	"api/event_db"
	"api/metrics_db"
	"api/models"
	"api/server"
	"fmt"
	"time"
)

// https://gin-gonic.com/docs/examples/custom-http-config/
// https://github.com/vsouza/go-gin-boilerplate
func main() {
	time.Sleep(10 * time.Second)
	currentTime := time.Now()
	fmt.Println("Current Time in String: ", currentTime.String())
	httpsRouter := server.NewRouter()
	httpRouter := server.NewRouter()
	err := metrics_db.Init()
	if err != nil {
		panic("No se pudo conectar con Metricas-DB")
	}
	err = event_db.Init()
	if err != nil {
		panic("No se pudo conectar con Event-DB")
	}
	tokens, err := models.RecoverTokens()
	fmt.Println(fmt.Sprintf("Se recuperaron %d tokens", tokens))
	if err != nil {
		panic("No se pudo recuperar los tokens")
	}
	go httpRouter.Run(":80")
	httpsRouter.RunTLS(":443", "localhost.crt", "localhost.key")
}
