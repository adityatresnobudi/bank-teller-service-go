package service

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/service_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type ServiceService interface {
	GetAll(ctx context.Context) (*dto.GetAllServicesResponseDTO, errs.MessageErr)
	GetOne(ctx context.Context, id string) (*dto.GetOneServiceResponseDTO, errs.MessageErr)
	Create(ctx context.Context, payload dto.CreateServiceRequestDTO) (*dto.CreateServiceResponseDTO, errs.MessageErr)
	UpdateById(ctx context.Context, id string, payload dto.UpdateServiceRequestDTO) (*dto.UpdateServiceResponseDTO, errs.MessageErr)
	DeleteById(ctx context.Context, id string) (*dto.DeleteByIdServiceResponseDTO, errs.MessageErr)
}

type serviceServiceIMPL struct {
	serviceRepo service_repo.Repository
}

func NewServiceService(serviceRepo service_repo.Repository) ServiceService {
	return &serviceServiceIMPL{
		serviceRepo: serviceRepo,
	}
}

func (s *serviceServiceIMPL) GetAll(ctx context.Context) (*dto.GetAllServicesResponseDTO, errs.MessageErr) {
	services, err := s.serviceRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	result := dto.GetAllServicesResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data:                  entity.Services(services).ToSliceOfServicesResponseDTO(),
	}

	return &result, nil
}

func (s *serviceServiceIMPL) GetOne(ctx context.Context, id string) (*dto.GetOneServiceResponseDTO, errs.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	service, err := s.serviceRepo.GetOneById(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	result := dto.GetOneServiceResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data:                  *service.ToServiceResponseDTO(),
	}

	return &result, nil
}

func (s *serviceServiceIMPL) Create(ctx context.Context, payload dto.CreateServiceRequestDTO) (*dto.CreateServiceResponseDTO, errs.MessageErr) {
	existingService, err := s.serviceRepo.GetOneByCode(
		ctx,
		payload.Code,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	if existingService != nil {
		return nil, errs.NewConflictError("service already exists")
	}

	service := entity.Service{
		Code: payload.Code,
		Name: payload.Name,
	}

	err = s.serviceRepo.Create(ctx, service)

	if err != nil {
		return nil, err
	}

	result := dto.CreateServiceResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Service created successfully"},
	}

	return &result, nil
}

func (s *serviceServiceIMPL) UpdateById(ctx context.Context, id string, payload dto.UpdateServiceRequestDTO) (*dto.UpdateServiceResponseDTO, errs.MessageErr) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	service, errData := s.serviceRepo.GetOneById(
		ctx,
		parsedId,
	)

	if errData != nil && errData.StatusCode() != http.StatusNotFound {
		return nil, errData
	}

	if service == nil {
		return nil, errs.NewBadRequest("service does not exist")
	}

	if payload.Code != "" {
		service.Code = payload.Code
	}

	if payload.Name != "" {
		service.Name = payload.Name
	}

	updatedService, errData := s.serviceRepo.UpdateById(ctx, *service)
	if errData != nil {
		return nil, errData
	}

	result := dto.UpdateServiceResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Service updated successfully"},
		Data:                  *updatedService.ToServiceResponseDTO(),
	}

	return &result, nil
}

func (s *serviceServiceIMPL) DeleteById(ctx context.Context, id string) (*dto.DeleteByIdServiceResponseDTO, errs.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	_, err := s.serviceRepo.GetOneById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	err = s.serviceRepo.DeleteById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	result := dto.DeleteByIdServiceResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Service deleted successfully"},
	}

	return &result, nil
}
