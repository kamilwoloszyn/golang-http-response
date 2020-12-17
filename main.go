package main

import (
	"net/http"

	"github.com/kamilwoloszyn/golang-http-response/routes"
)

func main() {
	http.HandleFunc("/", routes.Mainroute)
	http.HandleFunc("/personal", routes.Personal)
	http.HandleFunc("/visitcounter", routes.VisitCounter)
	http.ListenAndServe(":8080", nil)
}
