package models

import "time"

type Booking struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	TimeSlotID string    `json:"time_slot_id"`
	CreatedAt  time.Time `json:"created_at"`
}
