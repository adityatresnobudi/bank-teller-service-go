package service_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/service_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type servicePG struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) service_repo.Repository {
	return &servicePG{
		db: db,
	}
}

func (s *servicePG) GetAll(ctx context.Context) ([]entity.Service, errs.MessageErr) {
	rows, err := s.db.QueryContext(ctx, GET_ALL_SERVICE)

	if err != nil {
		log.Printf("db get all services: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	services := []entity.Service{}

	for rows.Next() {
		service := entity.Service{}

		if err = rows.Scan(
			&service.Id,
			&service.Code,
			&service.Name,
			&service.CreatedAt,
			&service.UpdatedAt,
		); err != nil {
			log.Printf("db scan get all services: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		services = append(services, service)
	}

	return services, nil
}
func (s *servicePG) GetOneById(ctx context.Context, id uuid.UUID) (*entity.Service, errs.MessageErr) {
	service := entity.Service{}

	if err := s.db.QueryRowContext(
		ctx,
		GET_SERVICE_BY_ID,
		id,
	).Scan(
		&service.Id,
		&service.Name,
		&service.Code,
		&service.CreatedAt,
		&service.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one service by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("service was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &service, nil
}
func (s *servicePG) GetOneByCode(ctx context.Context, code string) (*entity.Service, errs.MessageErr) {
	service := entity.Service{}

	if err := s.db.QueryRowContext(
		ctx,
		GET_SERVICE_BY_CODE,
		code,
	).Scan(
		&service.Id,
		&service.Name,
		&service.Code,
		&service.CreatedAt,
		&service.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one service by email: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("service was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &service, nil
}
func (s *servicePG) Create(ctx context.Context, service entity.Service) errs.MessageErr {
	if _, err := s.db.ExecContext(
		ctx,
		INSERT_SERVICE,
		service.Code,
		service.Name,
	); err != nil {
		log.Printf("db scan create service: %s\n", err.Error())
		return errs.NewInternalServerError()
	}

	return nil
}
func (s *servicePG) UpdateById(ctx context.Context, service entity.Service) (*entity.Service, errs.MessageErr) {
	updatedService := entity.Service{}

	if err := s.db.QueryRowContext(
		ctx,
		UPDATE_SERVICE,
		service.Code,
		service.Name,
		service.Id,
	).Scan(
		&updatedService.Id,
		&updatedService.Code,
		&updatedService.Name,
		&updatedService.CreatedAt,
		&updatedService.UpdatedAt,
	); err != nil {
		log.Printf("db scan update service by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("service was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &updatedService, nil
}
func (s *servicePG) DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr {
	if _, err := s.db.ExecContext(
		ctx,
		DELETE_SERVICE,
		id,
	); err != nil {
		log.Printf("db delete service by id: %s\n", err.Error())
		return errs.NewInternalServerError()
	}

	return nil
}
