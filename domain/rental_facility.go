package domain

import "time"

type RentalFacility struct {
	Id        *int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeleteAt  *time.Time
}
