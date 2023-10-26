package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Cwjiee/goproxy/routes"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	
	var customTransport = http.DefaultTransport

	targetURL, err := routes.CustomRoutes(r, w)
	if err != nil {
		log.Printf("Error routing request URL %v", err)
		http.Error(w, "Error routing request", http.StatusInternalServerError)
		return
	}

	log.Printf("Received API response %v %v", r.Method, r.URL)

	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		log.Printf("Error creating API response %v", err)
		http.Error(w, "There's an error creating the proxy request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	fmt.Println(proxyReq.Header)
	proxyReq.Header.Set("User-Agent", "Anonymous")
	fmt.Println(proxyReq.Header)

	res, err := customTransport.RoundTrip(proxyReq)

	if err != nil {
		log.Printf("Error sending API response %v", err)
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
