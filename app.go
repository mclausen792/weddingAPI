package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("deployed")
	r := mux.NewRouter()
	r.HandleFunc("/guests", GuestsEndpoint).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("running on Port" + port)
	http.ListenAndServe(":"+port, handlers.CORS()(r))
}

func GuestsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /guests POST:")
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
	r.Body.Close()
	w.WriteHeader(http.StatusOK)
}
