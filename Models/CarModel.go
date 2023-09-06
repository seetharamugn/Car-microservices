package Models

type Car struct {
	Id       string `json:"Id"`
	Make     string `json:"Make"`
	Model    string `json:"Model"`
	Package  string `json:"Package"`
	Color    string `json:"Color"`
	Year     int    `json:"Year"`
	Category string `json:"Category"`
	Mileage  int    `json:"Mileage (ml)"`
	Price    int    `json:"Price (cents)"`
}
