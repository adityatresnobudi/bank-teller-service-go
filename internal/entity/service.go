package entity

import (
	"time"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/google/uuid"
)

type Service struct {
	Id        uuid.UUID
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Services []Service

func (s *Service) ToServiceResponseDTO() *dto.ServiceResponseDTO {
	return &dto.ServiceResponseDTO{
		Id:        s.Id,
		Code:      s.Code,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func (s Services) ToSliceOfServicesResponseDTO() []dto.ServiceResponseDTO {
	result := []dto.ServiceResponseDTO{}
	for _, service := range s {
		result = append(result, *service.ToServiceResponseDTO())
	}

	return result
}
