package main

import (
	"log"
	"net/http"
)

func main() {
	
	server := http.Server {
		Addr: ":8081",
		Handler: http.HandlerFunc(handleRequest), 
	}

	log.Println("Starting proxy server on :8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy server", err)
	}
}
