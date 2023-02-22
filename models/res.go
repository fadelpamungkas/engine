package models

type Response struct {
	Status  int         `json:"status" validate:"required"`
	Message string      `json:"message" validate:"required"`
	Data    interface{} `json:"data" validate:"required"`
}
