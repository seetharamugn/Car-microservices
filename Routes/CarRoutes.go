package Routes

import (
	"fmt"
	"github.com/seetharamugn/car-microservices/Controllers"
	"net/http"
)

func SetupRotes() {
	http.HandleFunc("/CreateNewCar", Controllers.CreateNewCar)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
