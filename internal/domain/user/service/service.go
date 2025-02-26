package service

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/entity"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/user_repo"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
)

type UserService interface {
	GetAll(ctx context.Context) (*dto.GetAllUsersResponseDTO, errs.MessageErr)
	GetOne(ctx context.Context, id string) (*dto.GetOneUserResponseDTO, errs.MessageErr)
	Create(ctx context.Context, payload dto.CreateUserRequestDTO) (*dto.CreateUserResponseDTO, errs.MessageErr)
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

func (u *userServiceIMPL) GetOne(ctx context.Context, id string) (*dto.GetOneUserResponseDTO, errs.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	user, err := u.userRepo.GetOneById(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	result := dto.GetOneUserResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data:                  *user.ToUserResponseDTO(),
	}

	return &result, nil
}

func (u *userServiceIMPL) Create(ctx context.Context, payload dto.CreateUserRequestDTO) (*dto.CreateUserResponseDTO, errs.MessageErr) {
	if err := u.createValidator(payload); err != nil {
		return nil, err
	}

	existingUser, err := u.userRepo.GetOneByEmail(
		ctx,
		payload.Email,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	if existingUser != nil {
		return nil, errs.NewConflictError("user already exists")
	}

	user := entity.User{
		Email: payload.Email,
		Password: payload.Password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	err = u.userRepo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	result := dto.CreateUserResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "User created successfully"},
	}

	return &result, nil
}
