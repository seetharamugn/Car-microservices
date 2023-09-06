package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/seetharamugn/car-microservices/Models"
	"github.com/seetharamugn/car-microservices/Services"
	"net/http"
)

func CreateNewCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var newCar Models.Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := Services.CreateNewCar(newCar)
	fmt.Println(resp)
	carJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal car list to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carJSON)
}

func GetCarList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	resp, err := Services.GetCarList()
	if err != nil {

	}
	carListJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal car list to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carListJSON)
}
