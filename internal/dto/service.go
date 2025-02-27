package dto

import (
	"time"

	"github.com/google/uuid"
)

type ServiceResponseDTO struct {
	Id        uuid.UUID `json:"id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
	Code      string    `json:"code" example:"CGK"`
	Name      string    `json:"name" example:"adit"`
	CreatedAt time.Time `json:"created_at" example:"2025-02-22T15:11:19.25616+07:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-02-22T15:11:19.25616+07:00"`
} // @name ServiceResponse

type CreateServiceRequestDTO struct {
	Code    string `json:"code" example:"CGK"`
	Name string `json:"name" example:"service"`
} // @name CreateServiceRequest

type UpdateServiceRequestDTO struct {
	Code    string `json:"code" example:"CGK"`
	Name string `json:"name" example:"service"`
} // @name UpdateServiceRequest

type GetAllServicesResponseDTO struct {
	CommonBaseResponseDTO
	Data []ServiceResponseDTO
} // @name GetAllServicesResponse

type GetOneServiceResponseDTO struct {
	CommonBaseResponseDTO
	Data ServiceResponseDTO `json:"data"`
} // @name GetOneServiceResponse

type CreateServiceResponseDTO struct {
	CommonBaseResponseDTO
} // @name CreateServiceResponse

type UpdateServiceResponseDTO struct {
	CommonBaseResponseDTO
	Data ServiceResponseDTO `json:"data"`
} // @name UpdateServiceResponse

type DeleteByIdServiceResponseDTO struct {
	CommonBaseResponseDTO
} // @name DeleteServiceResponse
