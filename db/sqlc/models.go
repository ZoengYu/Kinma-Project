// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Fundraise struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	// must be positive
	TargetAmount int64 `json:"target_amount"`
	// must be positive
	ProgressAmount int64        `json:"progress_amount"`
	Success        bool         `json:"success"`
	StartDate      time.Time    `json:"start_date"`
	EndDate        sql.NullTime `json:"end_date"`
}

type Product struct {
	ID         int64        `json:"id"`
	AccountID  int64        `json:"account_id"`
	Title      string       `json:"title"`
	Content    string       `json:"content"`
	ProductTag []string     `json:"product_tag"`
	CreatedAt  time.Time    `json:"created_at"`
	LastUpdate sql.NullTime `json:"last_update"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToProductID   int64 `json:"to_product_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Success   bool      `json:"success"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}