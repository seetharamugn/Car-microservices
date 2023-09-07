package Repositories

import (
	"fmt"
	"github.com/seetharamugn/car-microservices/Models"
	"math/rand"
	"time"
)

// result is the in-memory to store the car details
var result = make(map[string]interface{})

// CreateNewCar is to Create a new car
func CreateNewCar(car Models.Car) interface{} {
	car.Id = GenerateRandomString(8)
	result[car.Id] = car
	fmt.Println("Car inserted successfully with ID:", car.Id)
	return result[car.Id]
}

// GetCarList is to get all the cars
func GetCarList() interface{} {
	return result
}

// GetCar is the Fetch particular car details
func GetCar(carId string) interface{} {
	return result[carId]
}

// UpdateCar is the Modify the details of the particular car
func UpdateCar(carId string, updatedCar Models.Car) interface{} {
	result[carId] = Models.Car{
		Id:       carId,
		Make:     updatedCar.Make,
		Model:    updatedCar.Model,
		Package:  updatedCar.Package,
		Color:    updatedCar.Color,
		Year:     updatedCar.Year,
		Category: updatedCar.Category,
		Mileage:  updatedCar.Mileage,
		Price:    updatedCar.Price,
	}
	return result[carId]
}

// GenerateRandomString is the 8 digit random number generator
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}
