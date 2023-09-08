package Controllers

import (
	"encoding/json"
	"github.com/seetharamugn/car-microservices/Dao"
	"github.com/seetharamugn/car-microservices/Models"
	"github.com/seetharamugn/car-microservices/Services"
	"log"
	"net/http"
)

// CreateNewCar crate new Car Details
func CreateNewCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Fatalf("Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var newCar Models.Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		log.Fatalf(err.Error(), http.StatusBadRequest)
		return
	}
	resp := Services.CreateNewCar(newCar)
	response := Dao.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       resp,
	}
	carJSON, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal car list to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carJSON)
	log.Println("Nea Car Details created successfully", http.StatusOK)
}

// GetCarList to Get all Car Details
func GetCarList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Fatalf("Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	resp := Services.GetCarList()
	response := Dao.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       resp,
	}
	carListJSON, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal car list to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carListJSON)
	log.Println("Fetch all car details successfully", http.StatusOK)
}

// GetCar is Particular Car using the CarId the car id is unique
func GetCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Fatalf("Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	carId := r.URL.Query().Get("id")

	resp := Services.GetCar(carId)

	response := Dao.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       resp,
	}
	carJSON, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal car list to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carJSON)
	log.Println("Fetch the car details successfully", http.StatusOK)
}

// UpdateCar update the particular car details
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Fatalf("Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	carId := r.URL.Query().Get("id")
	var updateCar Models.Car
	err := json.NewDecoder(r.Body).Decode(&updateCar)
	if err != nil {
		log.Fatalf(err.Error(), http.StatusBadRequest)
		return
	}
	resp := Services.UpdateCar(carId, updateCar)

	response := Dao.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       resp,
	}
	carJSON, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal car list to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carJSON)
	log.Println("updated the car details successfully", http.StatusOK)
}
