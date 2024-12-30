package domain

import "time"

type User struct {
	ID       uint      `json:"id"`
	FirtName string    `json:"first_name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Code     int       `json:"code"`
	Expiry   time.Time `json:"expiry"`
	Verified bool      `json:"verify"`
	UserType string    `json:"user_type"`
}
