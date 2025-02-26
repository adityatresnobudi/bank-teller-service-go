package entity

import (
	"time"

	"github.com/google/uuid"
)

type Queue struct {
	Id        uuid.UUID
	Status    string
	QueueNum  string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    uuid.UUID
	ServiceId uuid.UUID
}
