package Controllers

import (
	"encoding/json"
	"github.com/seetharamugn/car-microservices/Models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateNewCar(t *testing.T) {
	newCar := Models.Car{
		Make:     "FORD",
		Model:    "BRNCO",
		Package:  "Badiands",
		Color:    "BrutOrange",
		Year:     2022,
		Category: "SUV",
		Mileage:  1,
		Price:    449900,
	}
	newCarJSON, _ := json.Marshal(newCar)
	req := httptest.NewRequest("POST", "/create", strings.NewReader(string(newCarJSON)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	CreateNewCar(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}
}

func TestGetCar(t *testing.T) {
	req := httptest.NewRequest("GET", "/get?id=dRSTnYnR", nil)
	rr := httptest.NewRecorder()
	GetCar(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}
}

func TestGetCarList(t *testing.T) {
	req := httptest.NewRequest("GET", "/list", nil)
	rr := httptest.NewRecorder()
	GetCarList(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateCar(t *testing.T) {
	updateCar := Models.Car{
		Make:     "FORD",
		Model:    "BRNCO",
		Package:  "Badiands",
		Color:    "BrutOrange",
		Year:     2022,
		Category: "SUV",
		Mileage:  1,
		Price:    449900,
	}
	updateCarJSON, _ := json.Marshal(updateCar)
	req := httptest.NewRequest("PUT", "/update?id=dRSTnYnR", strings.NewReader(string(updateCarJSON)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	UpdateCar(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}
}
