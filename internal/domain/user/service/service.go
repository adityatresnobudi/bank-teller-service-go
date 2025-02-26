package service

import (
	"context"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/user_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
)

type UserService interface {
	GetAll(ctx context.Context) (*dto.GetAllUsersResponseDTO, errs.MessageErr)
}

type userServiceIMPL struct {
	userRepo user_repo.Repository
}

func NewUserService(userRepo user_repo.Repository) UserService {
	return &userServiceIMPL{
		userRepo: userRepo,
	}
}

func (u *userServiceIMPL) GetAll(ctx context.Context) (*dto.GetAllUsersResponseDTO, errs.MessageErr) {
	users, err := u.userRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	result := dto.GetAllUsersResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data:                  entity.Users(users).ToSliceOfUsersResponseDTO(),
	}

	return &result, nil
}
