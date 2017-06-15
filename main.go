package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	sub := r.PathPrefix("/api").Subrouter()
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}
	http.ListenAndServe(port, nil)
}
