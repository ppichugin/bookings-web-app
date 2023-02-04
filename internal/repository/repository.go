package repository

import "github.com/ppichugin/booking-for-breakfast/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
