package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func CustomRoutes(r *http.Request) (*url.URL, error) {

	route := strings.Split(r.URL.Path, "/")

	if route[1] != "get" {
		return nil, fmt.Errorf("request route not found")
	}

	if len(route) == 2 {
		return url.Parse("http://localhost:8080/todos")
	} else if len(route) == 3 {
		resRoute := "http://localhost:8080/todos/" + route[2]
		return url.Parse(resRoute)
	} else {
		return url.Parse("http://localhost:8080/todos")
	} 
}
