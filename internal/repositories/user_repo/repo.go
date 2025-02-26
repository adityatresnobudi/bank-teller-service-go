package user_repo

import (
	"context"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entity.User, errs.MessageErr)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.User, errs.MessageErr)
	Create(ctx context.Context, user entity.User) (*entity.User, errs.MessageErr)
	UpdateById(ctx context.Context, user entity.User) (*entity.User, errs.MessageErr)
	DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr
}
