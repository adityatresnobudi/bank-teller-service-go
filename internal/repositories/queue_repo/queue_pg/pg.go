package queue_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/queue_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type queuePG struct {
	db *sql.DB
}

func NewQueueRepo(db *sql.DB) queue_repo.Repository {
	return &queuePG{
		db: db,
	}
}

func (q *queuePG) GetAll(ctx context.Context) ([]entity.Queue, errs.MessageErr) {
	rows, err := q.db.QueryContext(ctx, GET_ALL_QUEUE)

	if err != nil {
		log.Printf("db get all queues: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	queues := []entity.Queue{}

	for rows.Next() {
		queue := entity.Queue{}

		if err = rows.Scan(
			&queue.Id,
			&queue.Status,
			&queue.QueueNum,
			&queue.CreatedAt,
			&queue.UpdatedAt,
			&queue.UserId,
			&queue.ServiceId,
		); err != nil {
			log.Printf("db scan get all queues: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		queues = append(queues, queue)
	}

	return queues, nil
}
func (q *queuePG) GetOneById(ctx context.Context, id uuid.UUID) (*entity.Queue, errs.MessageErr) {
	queue := entity.Queue{}

	if err := q.db.QueryRowContext(
		ctx,
		GET_QUEUE_BY_ID,
		id,
	).Scan(
		&queue.Id,
		&queue.Status,
		&queue.QueueNum,
		&queue.CreatedAt,
		&queue.UpdatedAt,
		&queue.UserId,
		&queue.ServiceId,
	); err != nil {
		log.Printf("db scan get one queue by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("queue was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &queue, nil
}
func (q *queuePG) GetOneByQueueNum(ctx context.Context, queueNum string) (*entity.Queue, errs.MessageErr) {
	queue := entity.Queue{}

	if err := q.db.QueryRowContext(
		ctx,
		GET_QUEUE_BY_QUEUE_NUMBER,
		queueNum,
	).Scan(
		&queue.Id,
		&queue.Status,
		&queue.QueueNum,
		&queue.CreatedAt,
		&queue.UpdatedAt,
		&queue.UserId,
		&queue.ServiceId,
	); err != nil {
		log.Printf("db scan get one queue by queue number: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("queue was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &queue, nil
}
func (q *queuePG) Create(ctx context.Context, queue entity.Queue) errs.MessageErr {
	if _, err := q.db.ExecContext(
		ctx,
		INSERT_QUEUE,
		queue.QueueNum,
		queue.UserId,
		queue.ServiceId,
	); err != nil {
		log.Printf("db scan create queue: %s\n", err.Error())
		return errs.NewInternalServerError()
	}

	return nil
}
func (q *queuePG) UpdateById(ctx context.Context, queue entity.Queue) (*entity.Queue, errs.MessageErr) {
	updatedQueue := entity.Queue{}

	if err := q.db.QueryRowContext(
		ctx,
		UPDATE_QUEUE,
		queue.Status,
		queue.Id,
	).Scan(
		&updatedQueue.Id,
		&updatedQueue.Status,
		&updatedQueue.QueueNum,
		&updatedQueue.CreatedAt,
		&updatedQueue.UpdatedAt,
		&updatedQueue.UserId,
		&updatedQueue.ServiceId,
	); err != nil {
		log.Printf("db scan update queue status by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("queue was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &updatedQueue, nil
}
func (q *queuePG) DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr {
	if _, err := q.db.ExecContext(
		ctx,
		DELETE_QUEUE,
		id,
	); err != nil {
		log.Printf("db delete queue by id: %s\n", err.Error())
		return errs.NewInternalServerError()
	}

	return nil
}
