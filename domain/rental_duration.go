package domain

import "time"

type RentalDuration struct {
	Id        *int
	Unit      string
	Amount    int8
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeleteAt  *time.Time
}
