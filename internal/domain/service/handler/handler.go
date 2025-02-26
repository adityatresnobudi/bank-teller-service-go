package handler

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/domain/service/service"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/gin-gonic/gin"
)

type serviceHandler struct {
	r       *gin.Engine
	ctx     context.Context
	service service.ServiceService
}

func NewServiceHandler(
	r *gin.Engine,
	ctx context.Context,
	service service.ServiceService,
) *serviceHandler {
	return &serviceHandler{
		r:       r,
		ctx:     ctx,
		service: service,
	}
}

// @Summary Get All Services
// @Tags services
// @Produce json
// @Success 200 {object}  GetAllServicesResponse
// @Router /services [get]
func (s *serviceHandler) GetAll(c *gin.Context) {
	result, err := s.service.GetAll(s.ctx)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get One Service By ID
// @Tags services
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object}  GetOneServiceResponse
// @Router /services/{id} [get]
func (s *serviceHandler) GetOne(c *gin.Context) {
	id := c.Param("id")

	result, err := s.service.GetOne(s.ctx, id)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create Service
// @Tags services
// @Accept json
// @Produce json
// @Param requestBody body CreateServiceRequest true "Request Body"
// @Success 201 {object} CreateServiceResponse
// @Router /services [post]
func (s *serviceHandler) Create(c *gin.Context) {
	payload := dto.CreateServiceRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, err := s.service.Create(s.ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// @Summary Update Service
// @Tags services
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Param requestBody body UpdateServiceRequest true "Request Body"
// @Success 200 {object} UpdateServiceResponse
// @Router /services [put]
func (s *serviceHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")
	payload := dto.UpdateServiceRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, errData := s.service.UpdateById(s.ctx, id, payload)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Delete Service
// @Tags services
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 204 {object} DeleteServiceResponse
// @Router /services [delete]
func (s *serviceHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	result, errData := s.service.DeleteById(s.ctx, id)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusNoContent, result)
}
