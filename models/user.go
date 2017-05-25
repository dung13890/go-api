package models

import ()

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users struct {
	Users []User `json:"user"`
	Meta  int    `json:"meta"`
}
