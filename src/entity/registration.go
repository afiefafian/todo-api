package entity

import (
	"context"
	"time"
)

// Registration struct entity
// Used in json, postgres, validation input
type Registration struct {
	ID           ID        `json:"id,omitempty" pg:"id,pk,default:uuid_generate_v4()"`
	FirstName    string    `json:"first_name,omitempty" validate:"required" pg:"first_name"`
	LastName     string    `json:"last_name,omitempty" validate:"required" pg:"last_name"`
	Email        string    `json:"email" validate:"required,email" pg:"email"`
	Password     string    `json:"password,omitempty" validate:"required" pg:"password"`
	Phone        string    `json:"phone,omitempty" validate:"required" pg:"phone"`
	IsRegistered string    `pg:"is_registered"`
	CreatedAt    time.Time `json:"created_at,omitempty" pg:"created_at,default:now()"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" pg:"updated_at,default:now()"`
	DeletedAt    time.Time `json:"deleted_at,omitempty" pg:"deleted_at,soft_delete"`
}

// UserRegistrationServices represent the user's services
type RegistrationServices interface {
	Register(context.Context, *Registration) error
	//ResendCode(ctx context.Context, email string) error
	//VerifyCode(ctx context.Context, email string, code int8) error
}

// UserRegistration represent the user's repository contract
type RegistrationRepository interface {
	Store(ctx context.Context, u *Registration) error
	// GetByEmail(ctx context.Context, email string) (Registration, error)
	// UpdateStatus(ctx context.Context, u *User, id string) (bool, error)
}
