package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Cwjiee/goproxy/middleware"
	"github.com/Cwjiee/goproxy/routes"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	
	var customTransport = http.DefaultTransport

	targetURL, err := routes.CustomRoutes(r)
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

	proxyReq.Header.Set("User-Agent", "Anonymous")	

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

	filteredBody, contentLength, err := middleware.FilterContent(res.Body)
	if err != nil {
		log.Printf("Error censoring response body %v", err)
		http.Error(w, "There's an error censoring the response", http.StatusInternalServerError)
	}

	fmt.Println(res.Request.UserAgent())

	w.Header().Set("Content-Length", fmt.Sprintf("%d",contentLength))

	w.WriteHeader(res.StatusCode)

	io.Copy(w, filteredBody)
}
