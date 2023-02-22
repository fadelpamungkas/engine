package models

import "time"

type Entry struct {
	Id        int64     `json:"id"`
	AccountId int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
