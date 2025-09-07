package user

import "time"

type User struct {
	Id string
	Name string
	Email Email
	CreatedAt time.Time
}