package entity

import (
	"time"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (u *User) HashPassword() errs.MessageErr {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

	if err != nil {
		return errs.NewInternalServerError()
	}

	u.Password = string(b)

	return nil
}

func (u *User) Compare(password string) errs.MessageErr {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	); err != nil {
		return errs.NewUnauthorizedError("invalid password")
	}

	return nil
}
