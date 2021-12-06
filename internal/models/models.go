package models

import "time"

// User
type User struct {
	ID string
	FistName string
	LastName string
	Email string
	Phone string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Reservation
type Reservation struct {
	ID string
	FirstName string
	LastName string
	Phone string
	Email string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	StartDate time.Time
	EndDate	time.Time
	RoomId string
	Room Room
}

// Room
type Room struct {
	ID string
	RoomName string
	CreatedAt time.Time
	UpdatedAt time.Time
}