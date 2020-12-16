package routes

import (
	"fmt"
	"net/http"
)

func Mainroute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "No content here")
}
