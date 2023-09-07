package Routes

import (
	"github.com/seetharamugn/car-microservices/Controllers"
	"net/http"
)

func SetupRotes() {
	http.HandleFunc("/CreateNewCar", Controllers.CreateNewCar)
	http.HandleFunc("/GetCarList", Controllers.GetCarList)
	http.HandleFunc("/GetCar", Controllers.GetCar)
	http.HandleFunc("/UpdateCar", Controllers.UpdateCar)
}
