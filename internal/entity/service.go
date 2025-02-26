package entity

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Id uuid.UUID
	Code string
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}