package Repositories

import (
	"database/sql"
	"fmt"
	"github.com/seetharamugn/car-microservices/Initializers"
	"github.com/seetharamugn/car-microservices/Models"
	"math/rand"
	"time"
)

var db = Initializers.ConnectDb()

func CreateNewCar(car Models.Car) (interface{}, error) {
	car.Id = GenerateRandomString(8)
	query := `
        INSERT INTO cars ( id,make, model, package, color, year, category, mileage, price)
        VALUES (?,?,?,?,?,?,?,?,?)
    `
	_, err := db.Exec(query, car.Id, car.Make, car.Model, car.Package, car.Color, car.Year, car.Category, car.Mileage, car.Price)
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

func GetCar(carId string) (interface{}, error) {
	query := `SELECT * FROM cars WHERE id = ?`
	row := db.QueryRow(query, carId)

	var car Models.Car
	err := row.Scan(&car.Id, &car.Make, &car.Model, &car.Package, &car.Color, &car.Year, &car.Category, &car.Mileage, &car.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return a custom error if the car is not found.
			return nil, fmt.Errorf("Car not found")
		}
		fmt.Println(err.Error())
		return nil, err
	}

	return &car, nil
}

func UpdateCar(carId string, updatedCar Models.Car) (interface{}, error) {
	query := `SELECT * FROM cars WHERE id = ?`
	row := db.QueryRow(query, carId)

	var existingCar Models.Car
	err := row.Scan(&existingCar.Id, &existingCar.Make, &existingCar.Model, &existingCar.Package, &existingCar.Color, &existingCar.Year, &existingCar.Category, &existingCar.Mileage, &existingCar.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Car not found")
		}
		fmt.Println(err.Error())
		return nil, err
	}

	updateQuery := `UPDATE cars SET make=?, model=?, package=?, color=?, year=?, category=?, mileage=?, price=? WHERE id = ?`
	_, updateErr := db.Exec(updateQuery, updatedCar.Make, updatedCar.Model, updatedCar.Package, updatedCar.Color, updatedCar.Year, updatedCar.Category, updatedCar.Mileage, updatedCar.Price, carId)
	if updateErr != nil {
		fmt.Println(updateErr.Error())
		return nil, updateErr
	}

	return updatedCar, nil
}
