package main

import (
	"api/event_db"
	"api/metrics_db"
	"api/models"
	"api/server"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	currentTime := time.Now()
	// Por la falta de certificados firmados...
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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
