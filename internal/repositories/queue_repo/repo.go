package queue_repo

import (
	"context"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entity.Queue, errs.MessageErr)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.Queue, errs.MessageErr)
	GetOneByQueueNum(ctx context.Context, queueNum string) (*entity.Queue, errs.MessageErr)
	Create(ctx context.Context, user entity.Queue) errs.MessageErr
	UpdateById(ctx context.Context, user entity.Queue) (*entity.Queue, errs.MessageErr)
	DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr
}
