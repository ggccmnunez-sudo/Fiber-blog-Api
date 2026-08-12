package models

import "time"

type User struct {
	// this is for Database tables
	//Good prcatice to use a variable required struct in go
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` //comparing the hashed with the empty string used "-" dapat text siya like password
	CreatedAt time.Time `json:"created_at"`
}
