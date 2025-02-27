package service

import (
	"context"
	"net/http"
	"strconv"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/queue_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/service_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/user_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type QueueService interface {
	GetAll(ctx context.Context) (*dto.GetAllQueuesResponseDTO, errs.MessageErr)
	GetOne(ctx context.Context, id string) (*dto.GetOneQueueResponseDTO, errs.MessageErr)
	Create(ctx context.Context, payload dto.CreateQueueRequestDTO) (*dto.CreateQueueResponseDTO, errs.MessageErr)
	UpdateByQueueNum(ctx context.Context, payload dto.UpdateQueueRequestDTO) (*dto.UpdateQueueResponseDTO, errs.MessageErr)
	DeleteById(ctx context.Context, id string) (*dto.DeleteByIdQueueResponseDTO, errs.MessageErr)
}

type queueServiceIMPL struct {
	queueRepo   queue_repo.Repository
	serviceRepo service_repo.Repository
	userRepo    user_repo.Repository
}

func NewQueueService(queueRepo queue_repo.Repository, serviceRepo service_repo.Repository, userRepo user_repo.Repository) QueueService {
	return &queueServiceIMPL{
		queueRepo:   queueRepo,
		serviceRepo: serviceRepo,
		userRepo:    userRepo,
	}
}

func (q *queueServiceIMPL) GetAll(ctx context.Context) (*dto.GetAllQueuesResponseDTO, errs.MessageErr) {
	queues, err := q.queueRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	result := dto.GetAllQueuesResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data:                  entity.Queues(queues).ToSliceOfQueuesResponseDTO(),
	}

	return &result, nil
}

func (q *queueServiceIMPL) GetOne(ctx context.Context, id string) (*dto.GetOneQueueResponseDTO, errs.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	queue, err := q.queueRepo.GetOneById(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	result := dto.GetOneQueueResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data:                  *queue.ToQueueResponseDTO(),
	}

	return &result, nil
}

func (q *queueServiceIMPL) Create(ctx context.Context, payload dto.CreateQueueRequestDTO) (*dto.CreateQueueResponseDTO, errs.MessageErr) {
	serviceCode := payload.ServiceName[:1]
	existingUser, err := q.userRepo.GetOneByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	existingService, err := q.serviceRepo.GetOneByName(ctx, payload.ServiceName)
	if err != nil {
		return nil, err
	}

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	queueNum, err := q.queueRepo.GetLatestQueueNum(ctx, existingService.Id)
	if err != nil {
		return nil, err
	}

	latestQueueNum, errData := strconv.Atoi(queueNum.QueueNum[1:])
	if errData != nil {
		return nil, errs.NewInternalServerError()
	}

	newQueueNum := latestQueueNum + 1
	qNString := "00" + strconv.Itoa(newQueueNum)
	switch len(qNString) {
	case 4:
		qNString = qNString[0:]
	case 5:
		qNString = qNString[1:]
	}
	qNString = serviceCode + qNString

	_, err = q.queueRepo.GetOneByQueueNum(ctx, qNString)
	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	newQueue := entity.Queue{
		QueueNum:  qNString,
		UserId:    existingUser.Id,
		ServiceId: existingService.Id,
	}

	err = q.queueRepo.Create(ctx, newQueue)

	if err != nil {
		return nil, err
	}

	result := dto.CreateQueueResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Queue created successfully"},
	}

	return &result, nil
}

func (q *queueServiceIMPL) UpdateByQueueNum(ctx context.Context, payload dto.UpdateQueueRequestDTO) (*dto.UpdateQueueResponseDTO, errs.MessageErr) {
	queue, errData := q.queueRepo.GetOneByQueueNum(
		ctx,
		payload.QueueNum,
	)

	if errData != nil && errData.StatusCode() != http.StatusNotFound {
		return nil, errData
	}

	if queue == nil {
		return nil, errs.NewBadRequest("queue does not exist")
	}

	switch queue.Status {
	case "pending":
		queue.Status = "processed"
	case "processed":
		queue.Status = "completed"
	}

	updatedQueue, errData := q.queueRepo.UpdateByQueueNum(ctx, *queue)
	if errData != nil {
		return nil, errData
	}

	result := dto.UpdateQueueResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Queue updated successfully"},
		Data:                  *updatedQueue.ToQueueResponseDTO(),
	}

	return &result, nil
}

func (q *queueServiceIMPL) DeleteById(ctx context.Context, id string) (*dto.DeleteByIdQueueResponseDTO, errs.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	_, err := q.queueRepo.GetOneById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	err = q.queueRepo.DeleteById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	result := dto.DeleteByIdQueueResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Queue deleted successfully"},
	}

	return &result, nil
}
