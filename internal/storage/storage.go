package storage

import (
	"time"

	"example.com/internal/models"
)

type Storage interface {
	CreateUser(email string, pwd string, username string, address string) (string, error)
	GetUserById(userID string) (*models.User, error)
	DeleteUserByID(userID string) error
	CreateTimeSlot(starTime time.Time, endTime time.Time) (string, error)
	GetTimeSlotByID(timeSlotID string) (*models.TimeSlot, error)
	DeleteTimeSlotByID(timeSlotID string) error
	IsTimeSlotBooked(timeSlotID string) (bool, error)
	CreateBooking(userID, timeSlotID string) (string, error)
	GetBookingById(bookingID string) (*models.Booking, error)
	DeleteBookingByID(bookingID string) error
}
