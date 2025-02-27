package entity

import (
	"time"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
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

type Queues []Queue

func (q *Queue) ToQueueResponseDTO() *dto.QueueResponseDTO {
	return &dto.QueueResponseDTO{
		Id:        q.Id,
		Status:    q.Status,
		QueueNum:  q.QueueNum,
		CreatedAt: q.CreatedAt,
		UpdatedAt: q.UpdatedAt,
		UserId:    q.UserId,
		ServiceId: q.ServiceId,
	}
}

func (q Queues) ToSliceOfQueuesResponseDTO() []dto.QueueResponseDTO {
	result := []dto.QueueResponseDTO{}
	for _, queue := range q {
		result = append(result, *queue.ToQueueResponseDTO())
	}

	return result
}
