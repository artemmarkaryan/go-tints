package controller

import (
	"log"
	"net/http"
	"time"
)

func Serve() {
	server := http.Server{
		Addr:              ":80",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    1 << 10,
	}

	ConfigureRoutes()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
