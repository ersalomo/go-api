package db

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   string    `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}
type Entry struct {
	ID        int64     `json:"id"`
	AccountID string    `json:"owner"`
	Amount    string    `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Transfer struct {
	ID            int64     `json:"id"`
	FromAccountID string    `json:"owner"`
	ToAccountID   string    `json:"to_account_id"`
	Amount        string    `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
