package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID
	Name        string
	PhoneNumber string
	Password    string
	Role        string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
