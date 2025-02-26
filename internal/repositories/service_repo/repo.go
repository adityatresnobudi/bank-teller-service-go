package service_repo

import (
	"context"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entity.Service, errs.MessageErr)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.Service, errs.MessageErr)
	GetOneByCode(ctx context.Context, code string) (*entity.Service, errs.MessageErr)
	Create(ctx context.Context, user entity.Service) errs.MessageErr
	UpdateById(ctx context.Context, user entity.Service) (*entity.Service, errs.MessageErr)
	DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr
}
