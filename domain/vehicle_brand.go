package domain

import "time"

type VehicleBrand struct {
	Id        *int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeleteAt  *time.Time
}
