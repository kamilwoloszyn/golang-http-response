package routes

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var visited int64

func VisitCounter(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	atomic.AddInt64(&visited, 1)
	fmt.Fprintf(w, "Visited during this session %d", visited)
}
