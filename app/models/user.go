package models

type User struct {
	IdNumber    string
	Email       string
	PhoneNumber string
	Address     string
	FirstName   string
	LastName    string
	Provider    string
	ProviderId  int8
	PhotoIdCard string
}
