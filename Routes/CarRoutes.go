package Routes

import (
	"github.com/seetharamugn/car-microservices/Controllers"
	"net/http"
)

func SetupRotes() {
	http.HandleFunc("/CreateNewCar", Controllers.CreateNewCar)
	http.HandleFunc("/GetCarList", Controllers.GetCarList)

}
