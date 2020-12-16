package routes

import (
	"fmt"
	"net/http"
)

func VisitCounter(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "It works!")
}
