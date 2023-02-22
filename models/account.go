package models

import "time"

type Account struct {
	Id         int64     `json:"id" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Balance    int64     `json:"balance" validate:"required"`
	Currency   string    `json:"currency" validate:"required"`
	Created_at time.Time `json:"created_at" validate:"required"`
}

type AccountRequest struct {
	Name     string `json:"name" validate:"required"`
	Balance  int64  `json:"balance" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type AccountResponse struct {
	Name     string `json:"name" validate:"required"`
	Balance  int64  `json:"balance" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}
