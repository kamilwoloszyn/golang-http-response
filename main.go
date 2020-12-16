package main

import (
	"net/http"

	"github.com/kamilwoloszyn/golang-http-response/routes"
)

var visited int64

func main() {
	http.HandleFunc("/personal", routes.Personal)
	http.HandleFunc("/visitcounter", routes.VisitCounter)
	http.ListenAndServe(":8080", nil)
}
