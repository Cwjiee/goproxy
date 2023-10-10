package main

import (
	"io"
	"log"
	"net/http"
)

var customTransport = http.DefaultClient.Transport

func handleRequest(w http.ResponseWriter, r *http.Request) {
	
	targetURL := r.URL
	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)

	if err != nil {
		http.Error(w, "There's an error creating the proxy request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	res, err := customTransport.RoundTrip(proxyReq)

	if err != nil {
		http.Error(w, "There's an error sending proxy request", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	for name, values := range res.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(res.StatusCode)

	io.Copy(w, res.Body)
}

func main() {
	
	server := http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(handleRequest), 
	}

	log.Println("Starting proxy server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy server", err)
	}
}
