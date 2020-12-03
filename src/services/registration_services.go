package services

import (
	"context"
	"errors"
	"strings"
	"todo_api/src/entity"
)

type registrationServices struct {
	userRepo         entity.UserRepository
	registrationRepo entity.RegistrationRepository
}

// NewRegistrationServices create new register services
func NewRegistrationServices(u entity.UserRepository, r entity.RegistrationRepository) entity.RegistrationServices {
	return &registrationServices{
		userRepo:         u,
		registrationRepo: r,
	}
}

// Register and send register otp to user email
func (r *registrationServices) Register(ctx context.Context, registration *entity.Registration) error {
	// Check email in db
	email := strings.Trim(registration.Email, "")
	user, err := r.userRepo.GetByEmail(ctx, email)
	if err != nil && err.Error() != "pg: no rows in result set" {
		return err
	}

	if user != (entity.User{}) {
		return errors.New("Email already used")
	}
	//if user
	// If found then return error
	return nil
}
