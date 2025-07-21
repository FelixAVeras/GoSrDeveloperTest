package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	ID     int    `json:"id"`
	Marca  string `json:"marca"`
	Modelo string `json:"modelo"`
	Año    int    `json:"año"`
	Color  string `json:"color"`
	Placa  string `json:"placa"`
}

var vehicles = make(map[int]Vehicle)
var idCounter = 1
var mutex = &sync.Mutex{}

func getVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	defer mutex.Unlock()

	vehicleList := make([]Vehicle, 0, len(vehicles))

	for _, vehicle := range vehicles {
		vehicleList = append(vehicleList, vehicle)
	}

	json.NewEncoder(w).Encode(vehicleList)
}

func createVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	defer mutex.Unlock()

	var vehicle Vehicle

	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vehicle.ID = idCounter
	idCounter++
	vehicles[vehicle.ID] = vehicle

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehicle)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/vehicles", getVehicles).Methods("GET")
	router.HandleFunc("/vehicles", createVehicle).Methods("POST")

	log.Println("API de vehículos escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
