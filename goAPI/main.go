package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func helloHandlers(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hola desde Go API!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}