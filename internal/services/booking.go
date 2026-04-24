package services

import (
	"fmt"
)

func (s *Service) CreateBooking(userId string, slotId string) (string, error) {
	_, err := s.repo.GetUserById(userId)
	if err != nil {
		return "", fmt.Errorf("user with user ID: %d not found", userId)
	}

	_, err = s.repo.GetTimeSlotByID(slotId)
	if err != nil {
		return "", fmt.Errorf("slot with slot ID: %d not found", slotId)
	}

	isBooked, err := s.repo.IsTimeSlotBooked(slotId)
	if isBooked {
		return "", fmt.Errorf("slot is already booked")
	}

	id, err = s.repo.CreateBooking(userId, slotId)
	return id, nil
}
