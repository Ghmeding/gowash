package storage

import (
	"errors"
	_ "fmt"
	"time"

	"example.com/internal/models"
	"github.com/google/uuid"
)

type MemoryRepo struct {
	users     map[string]*models.User
	timeSlots map[string]*models.TimeSlot
	bookings  map[string]*models.Booking
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		users:     make(map[string]*models.User),
		timeSlots: make(map[string]*models.TimeSlot),
		bookings:  make(map[string]*models.Booking),
	}
}

func (r *MemoryRepo) CreateUser(email string, pwd string, username string, address string) (string, error) {
	id := uuid.New().String()
	user := &models.User{
		ID:       id,
		Email:    email,
		Pwd:      pwd,
		Username: username,
		Address:  address,
	}

	r.users[id] = user
	return user.ID, nil
}

func (r *MemoryRepo) GetUserByID(id string) (*models.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *MemoryRepo) DeleteUserByID(userID string) error {
	_, exists := r.users[userID]
	if !exists {
		return errors.New("user not found")
	}
	delete(r.users, userID)
	return nil
}

func (r *MemoryRepo) CreateTimeSlot(starTime time.Time, endTime time.Time) (string, error) {
	id := uuid.New().String()
	timeSlot := &models.TimeSlot{
		ID:        id,
		StartTime: starTime,
		EndTime:   endTime,
	}

	r.timeSlots[id] = timeSlot
	return timeSlot.ID, nil
}

func (r *MemoryRepo) GetTimeSlotByID(id string) (*models.TimeSlot, error) {
	timeSlot, exists := r.timeSlots[id]
	if !exists {
		return nil, errors.New("time slot not found")
	}
	return timeSlot, nil
}

func (r *MemoryRepo) DeleteTimeSlotByID(timeSlotID string) error {
	_, exists := r.timeSlots[timeSlotID]
	if !exists {
		return errors.New("time slot not found")
	}
	delete(r.timeSlots, timeSlotID)
	return nil
}

func (r *MemoryRepo) IsTimeSlotBooked(timeSlotID string) (bool, error) {
	for _, b := range r.bookings {
		if b.TimeSlotID == timeSlotID {
			return true, nil
		}
	}
	return false, nil
}

func (r *MemoryRepo) CreateBooking(userID, timeSlotID string) (string, error) {
	id := uuid.New().String()
	booking := &models.Booking{
		ID:         id,
		UserID:     userID,
		TimeSlotID: timeSlotID,
		CreatedAt:  time.Now(),
	}
	r.bookings[id] = booking
	return id, nil
}

func (r *MemoryRepo) GetBookingByID(id string) (*models.Booking, error) {
	booking, exists := r.bookings[id]
	if !exists {
		return nil, errors.New("time slot not found")
	}
	return booking, nil
}

func (r *MemoryRepo) DeleteBookingByID(bookingID string) error {
	_, exists := r.bookings[bookingID]
	if !exists {
		return errors.New("booking not found")
	}
	delete(r.timeSlots, bookingID)
	return nil
}
