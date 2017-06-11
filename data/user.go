package data

import "time"

type User struct {
	Id       int
	Uuid     string
	Name     string
	Email    string
	Password string
	CreatAt  time.Time
}

type Session struct{
	Id int
	Uuid string
	Email string
	UserId int
	CreatAt time.Time
}