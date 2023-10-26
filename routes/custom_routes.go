package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func CustomRoutes(r *http.Request, w http.ResponseWriter) (*url.URL, error) {

	route := strings.Split(r.URL.Path, "/")

	if len(route) == 2 {
		return url.Parse("http://localhost:8080/todos")
	} else if len(route) == 3 {
		resRoute := "http://localhost:8080/todos/" + route[2]
		return url.Parse(resRoute)
	} else {
		http.Error(w, "wrong request route", http.StatusInternalServerError)
		return nil, fmt.Errorf("request route not found")
	}
}
