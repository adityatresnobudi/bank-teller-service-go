package handler

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/domain/user/service"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	r       *gin.Engine
	ctx     context.Context
	service service.UserService
}

func NewUserHandler(
	r *gin.Engine,
	ctx context.Context,
	service service.UserService,
) *userHandler {
	return &userHandler{
		r:       r,
		ctx:     ctx,
		service: service,
	}
}

// @Summary Get All Users
// @Tags users
// @Produce json
// @Success 200 {object}  GetAllUsersResponse
// @Router /users [get]
func (u *userHandler) GetAll(c *gin.Context) {
	result, err := u.service.GetAll(u.ctx)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get One User By ID
// @Tags users
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object}  GetOneUserResponse
// @Router /users/{id} [get]
func (u *userHandler) GetOne(c *gin.Context) {
	id := c.Param("id")

	result, err := u.service.GetOne(u.ctx, id)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create User
// @Tags users
// @Accept json
// @Produce json
// @Param requestBody body CreateUserRequest true "Request Body"
// @Success 201 {object} CreateUserResponse
// @Router /users [post]
func (u *userHandler) Create(c *gin.Context) {
	payload := dto.CreateUserRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, err := u.service.Create(u.ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// @Summary Update User
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param requestBody body UpdateUserRequest true "Request Body"
// @Success 200 {object} UpdateUserResponse
// @Router /users [put]
func (u *userHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")
	payload := dto.UpdateUserRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, errData := u.service.UpdateById(u.ctx, id, payload)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Delete User
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} DeleteUserResponse
// @Router /users [delete]
func (u *userHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	result, errData := u.service.DeleteById(u.ctx, id)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusNoContent, result)
}