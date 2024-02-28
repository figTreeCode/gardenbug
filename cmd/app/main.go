package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")
    router.HandleFunc("/api/data", apiHandler).Methods("GET")

    // Start the server
    port := 8080
    fmt.Printf("Server is running on http://localhost:%d\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Home Page!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    // Your API logic here
    fmt.Fprintln(w, "API Endpoint - Provide your data here")
}
