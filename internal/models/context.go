package models

import "time"

type ContextUser struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Email   string    `json:"email"`
	Expire  time.Time `json:"exp"`
}
