package model

import "time"

type User struct {
	Id int
	Email string
	Password string
	CreatedAt time.Time
}
