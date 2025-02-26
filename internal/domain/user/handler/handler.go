package handler

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/domain/user/service"
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
