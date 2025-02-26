package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	Id          uuid.UUID `json:"id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
	Name        string    `json:"name" example:"adit"`
	PhoneNumber string    `json:"phone_number" example:"adit"`
	Password    string    `json:"password" example:"a;dkqenrpokldafj;akjdga"`
	Role        string    `json:"role" example:"customer"`
	Email       string    `json:"email" example:"adit@abc.xyz"`
	CreatedAt   time.Time `json:"created_at" example:"2025-02-22T15:11:19.25616+07:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-02-22T15:11:19.25616+07:00"`
} // @name UserResponse

type CreateUserRequestDTO struct {
	Email       string    `json:"email" example:"adit@abc.xyz"`
	Password    string    `json:"password" example:"a;dkqenrpokldafj;akjdga"`
} // @name CreateUserRequest

type UpdateUserRequestDTO struct {
	Name        string    `json:"name" example:"adit"`
	PhoneNumber string    `json:"phone_number" example:"adit"`
	Password    string    `json:"password" example:"a;dkqenrpokldafj;akjdga"`
	Role        string    `json:"role" example:"customer"`
	Email       string    `json:"email" example:"adit@abc.xyz"`
} // @name UpdateUserRequest

type GetAllUsersResponseDTO struct {
	CommonBaseResponseDTO
	Data []UserResponseDTO
} // @name GetAllUsersResponse

type GetOneUserResponseDTO struct {
	CommonBaseResponseDTO
	Data UserResponseDTO `json:"data"`
} // @name GetOneUserResponse

type CreateUserResponseDTO struct {
	CommonBaseResponseDTO
} // @name CreateUserResponse

type UpdateUserResponseDTO struct {
	CommonBaseResponseDTO
	Data UserResponseDTO `json:"data"`
} // @name UpdateUserResponse

type DeleteByIdUserResponseDTO struct {
	CommonBaseResponseDTO
} // @name DeleteUserResponse
