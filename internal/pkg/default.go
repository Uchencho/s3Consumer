package pkg

import (
	"github.com/Uchencho/commons/ctime"
	"github.com/Uchencho/commons/uuid"
)

// User is a representation of an addyflex user
type User struct {
	ID             uuid.V4     `json:"id"`
	Email          string      `json:"email" validate:"required,email"`
	HashedPassword string      `json:"password,omitempty"`
	FirstName      string      `json:"first_name"`
	PhoneNumber    string      `json:"phone_number"`
	UserAddress    string      `json:"user_address"`
	IsActive       bool        `json:"is_active"`
	DateJoined     ctime.Epoch `json:"date_joined"`
	LastLogin      ctime.Epoch `json:"last_login"`
	Longitude      string      `json:"longitude,omitempty"`
	Latitude       string      `json:"latitude,omitempty"`
	DeviceID       string      `json:"device_id,omitempty"`
}
