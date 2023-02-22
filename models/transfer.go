package models

import "time"

type Transfer struct {
	Id            int64     `json:"id"`
	FromAccountId int64     `json:"from_account_id"`
	ToAccountId   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
