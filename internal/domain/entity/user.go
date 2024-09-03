package entity

import (
	"time"
)

type Signin struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password," db:"password_hash"`
	AppID    int    `json:"app_id"`
}

type Signup struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password," db:"password_hash"`
}

type User struct {
	ID           string    `json:"id,omitempty" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash []byte    `json:"password," db:"password_hash"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	IsAdmin      bool      `json:"is_admin" db:"is_admin"`
}
