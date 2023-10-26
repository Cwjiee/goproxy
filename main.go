package main

import (
	"log"
	"net/http"

	"github.com/Cwjiee/goproxy/middleware"
)

func main() {
	middleware.InitLogger()
	
	server := http.Server {
		Addr: ":8081",
		Handler: http.HandlerFunc(HandleRequest), 
	}

	log.Println("Starting proxy server on :8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy server", err)
	}
}
