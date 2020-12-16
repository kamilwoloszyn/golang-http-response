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
		"Kamil",
		"Woloszyn",
	}
	v, err := json.Marshal(&respond)
	if err == nil {
		fmt.Sprintf("%s", string(v))
	} else {
		fmt.Sprintf("500 Internal server error!")
	}

}
