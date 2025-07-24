package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello, JSON!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server starting on :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
