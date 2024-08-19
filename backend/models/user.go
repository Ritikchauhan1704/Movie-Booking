package models

type USER struct {
	Id    int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}