package handler

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/domain/queue/service"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/gin-gonic/gin"
)

type queueHandler struct {
	r       *gin.Engine
	ctx     context.Context
	service service.QueueService
}

func NewQueueHandler(
	r *gin.Engine,
	ctx context.Context,
	service service.QueueService,
) *queueHandler {
	return &queueHandler{
		r:       r,
		ctx:     ctx,
		service: service,
	}
}

func (q *queueHandler) GetAll(c *gin.Context) {
	result, err := q.service.GetAll(q.ctx)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (q *queueHandler) GetOne(c *gin.Context) {
	id := c.Param("id")

	result, err := q.service.GetOne(q.ctx, id)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (q *queueHandler) Create(c *gin.Context) {
	payload := dto.CreateQueueRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, err := q.service.Create(q.ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (q *queueHandler) UpdateByQueueNum(c *gin.Context) {
	payload := dto.UpdateQueueRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, errData := q.service.UpdateByQueueNum(q.ctx, payload)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (q *queueHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	result, errData := q.service.DeleteById(q.ctx, id)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusNoContent, result)
}
