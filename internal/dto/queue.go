package dto

import (
	"time"

	"github.com/google/uuid"
)

type QueueResponseDTO struct {
	Id        uuid.UUID `json:"id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
	Status    string    `json:"status" example:"pending"`
	QueueNum  string    `json:"queue_number" example:"L001"`
	CreatedAt time.Time `json:"created_at" example:"2025-02-22T15:11:19.25616+07:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-02-22T15:11:19.25616+07:00"`
	UserId    uuid.UUID `json:"user_id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
	ServiceId uuid.UUID `json:"service_id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
} // @name QueueResponse

type CreateQueueRequestDTO struct {
	Email    string `json:"email" example:"adit@abc.xyz"`
	ServiceName string `json:"service_name" example:"loan"`
} // @name CreateQueueRequest

type UpdateQueueRequestDTO struct {
	QueueNum string `json:"queue_number" example:"L001"`
} // @name UpdateQueueRequest

type GetAllQueuesResponseDTO struct {
	CommonBaseResponseDTO
	Data []QueueResponseDTO
} // @name GetAllQueuesResponse

type GetOneQueueResponseDTO struct {
	CommonBaseResponseDTO
	Data QueueResponseDTO `json:"data"`
} // @name GetOneQueueResponse

type CreateQueueResponseDTO struct {
	CommonBaseResponseDTO
} // @name CreateQueueResponse

type UpdateQueueResponseDTO struct {
	CommonBaseResponseDTO
	Data QueueResponseDTO `json:"data"`
} // @name UpdateQueueResponse

type DeleteByIdQueueResponseDTO struct {
	CommonBaseResponseDTO
} // @name DeleteQueueResponse
