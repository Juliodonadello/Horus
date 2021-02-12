package server

import (
	"net/http"
	"time"
)

func Init() *http.Server {
	router := NewRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}
