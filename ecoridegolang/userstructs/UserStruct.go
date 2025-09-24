package userstructs

import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type Session struct {
	Name string `json:"username"`
	Uuid string `json:"uuid"`
	Expiry time.Time `json:"expiry"`
}

type Credentials struct {
	Username string `json:"name"`
	Password string `json:"password"`
}