package service

import (
	"regexp"

	"github.com/adityatresnobudi/bank-teller-service-go/internal/dto"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
)

func (u *userServiceIMPL) createValidator(payload dto.CreateUserRequestDTO) errs.MessageErr {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(payload.Email) {
		return errs.NewBadRequest("email is not valid")
	}
	return nil
}
