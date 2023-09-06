package Routes

import (
	"fmt"
	"net/http"
)

func SetupRotes() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
