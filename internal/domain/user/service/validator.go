package service

import (
	"regexp"

	"github.com/adityatresnobudi/bank-teller-service-go/pkg/errs"
)

func (u *userServiceIMPL) emailValidator(email string) errs.MessageErr {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return errs.NewBadRequest("email is not valid")
	}
	return nil
}
