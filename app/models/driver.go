package models

import "time"

type Driver struct {
	IdNumber        string
	DriverLicenseId string
	FirstName       string
	LastName        string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeleteAt        *time.Time
}
