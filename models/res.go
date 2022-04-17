package models

type Response struct {
	Message     string       `json:"message"`
	Restaurants []Restaurant `json:"restaurants"`
}
