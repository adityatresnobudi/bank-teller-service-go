package service

import (
	"context"
	"fmt"
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
	UpdateById(ctx context.Context, id string, payload dto.UpdateUserRequestDTO) (*dto.UpdateUserResponseDTO, errs.MessageErr)
	DeleteById(ctx context.Context, id string) (*dto.DeleteByIdUserResponseDTO, errs.MessageErr)
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
	if err := u.emailValidator(payload.Email); err != nil {
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
		Email:    payload.Email,
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

func (u *userServiceIMPL) UpdateById(ctx context.Context, id string, payload dto.UpdateUserRequestDTO) (*dto.UpdateUserResponseDTO, errs.MessageErr) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	if payload.Email != "" {
		if err := u.emailValidator(payload.Email); err != nil {
			return nil, err
		}
	}

	user, errData := u.userRepo.GetOneById(
		ctx,
		parsedId,
	)

	if errData != nil && errData.StatusCode() != http.StatusNotFound {
		return nil, errData
	}

	if user == nil {
		return nil, errs.NewBadRequest("user does not exist")
	}

	if payload.Name != "" {
		user.Name = payload.Name
	}
	fmt.Println(user.Name)

	if payload.PhoneNumber != "" {
		user.PhoneNumber = payload.PhoneNumber
	}

	if payload.Password != "" {
		user.Password = payload.Password
	}

	if payload.Email != "" {
		user.Email = payload.Email
	}

	updatedUser, errData := u.userRepo.UpdateById(ctx, *user)
	if errData != nil {
		return nil, errData
	}

	result := dto.UpdateUserResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "User updated successfully"},
		Data:                  *updatedUser.ToUserResponseDTO(),
	}

	return &result, nil
}

func (u *userServiceIMPL) DeleteById(ctx context.Context, id string) (*dto.DeleteByIdUserResponseDTO, errs.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	_, err := u.userRepo.GetOneById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.DeleteById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	result := dto.DeleteByIdUserResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Account deleted successfully"},
	}

	return &result, nil
}