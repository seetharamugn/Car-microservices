package Repositories

import (
	"fmt"
	"github.com/seetharamugn/car-microservices/Initializers"
	"github.com/seetharamugn/car-microservices/Models"
	"math/rand"
	"time"
)

var db = Initializers.ConnectDb()

func CreateNewCar(car Models.Car) (interface{}, error) {
	query := `
        INSERT INTO cars ( make, model, package, color, year, category, mileage, price)
        VALUES (?,?,?,?,?,?,?,?)
    `
	_, err := db.Exec(query, car.Make, car.Model, car.Package, car.Color, car.Year, car.Category, car.Mileage, car.Price)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Car inserted successfully with ID:", car.Id)
	return car, nil
}

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

func GetCarList() (interface{}, error) {
	query := `SELECT * FROM cars`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	cars := []Models.Car{}
	for rows.Next() {
		var car Models.Car
		err := rows.Scan(
			&car.Id,
			&car.Make,
			&car.Model,
			&car.Package,
			&car.Color,
			&car.Year,
			&car.Category,
			&car.Mileage,
			&car.Price,
		)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return cars, nil
}
