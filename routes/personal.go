package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kamilwoloszyn/golang-http-response/data"
)

func Personal(w http.ResponseWriter, r *http.Request) {
	//simulate respond
	respond := data.Data{
		Name:    "Kamil",
		Surname: "Woloszyn",
	}
	v, err := json.Marshal(&respond)
	if err == nil {
		fmt.Fprint(w, string(v))
	} else {
		fmt.Fprintf(w, "500 Internal server error!")
	}

}
