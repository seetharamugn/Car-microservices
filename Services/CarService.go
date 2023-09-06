package Services

import (
	"github.com/seetharamugn/car-microservices/Models"
	"github.com/seetharamugn/car-microservices/Repositories"
)

func CreateNewCar(car Models.Car) (interface{}, error) {
	return Repositories.CreateNewCar(car)
}

func GetCarList() (interface{}, error) {
	return Repositories.GetCarList()
}

func GetCar(carId string) (interface{}, error) {
	return Repositories.GetCar(carId)
}

func UpdateCar(carId string, car Models.Car) (interface{}, error) {
	return Repositories.UpdateCar(carId, car)

}
