package Dao

// Response struct for the JSON Response of the Request
type Response struct {
	StatusCode int         `json:"StatusCode"`
	Message    string      `json:"Message"`
	Data       interface{} `json:"Data"`
}
