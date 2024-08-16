package core

import "time"

type Account struct {
	ID              string    `json:"userId"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	GoogleAccountID *string   `json:"googleAccountId"`
	Password        *string   `json:"password"`
	IsVerified      bool      `json:"isVerified"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
