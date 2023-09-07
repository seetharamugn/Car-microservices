package Services

import (
	"github.com/seetharamugn/car-microservices/Models"
	"github.com/seetharamugn/car-microservices/Repositories"
)

func CreateNewCar(car Models.Car) interface{} {
	return Repositories.CreateNewCar(car)
}

func GetCarList() interface{} {
	return Repositories.GetCarList()
}

func GetCar(carId string) interface{} {
	return Repositories.GetCar(carId)
}

func UpdateCar(carId string, car Models.Car) interface{} {
	return Repositories.UpdateCar(carId, car)
}
