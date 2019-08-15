package model

type Employee struct {
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Company *Company `json:"company"`
}