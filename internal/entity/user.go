package entity

import (
	"time"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID
	Name        string
	PhoneNumber string
	Password    string
	Role        string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Users []User

func (u *User) ToUserResponseDTO() *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		Id:          u.Id,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
		Role:        u.Role,
		Email:       u.Email,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func (u Users) ToSliceOfUsersResponseDTO() []dto.UserResponseDTO {
	result := []dto.UserResponseDTO{}
	for _, user := range u {
		result = append(result, *user.ToUserResponseDTO())
	}

	return result
}
