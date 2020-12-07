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
	Email        string    `json:"email" validate:"required,email" pg:"email,unique"`
	Password     string    `json:"password,omitempty" validate:"required" pg:"password"`
	Phone        string    `json:"phone,omitempty" validate:"required" pg:"phone"`
	IsRegistered bool      `pg:"is_registered,default:false"`
	CreatedAt    time.Time `json:"created_at,omitempty" pg:"created_at,default:now()"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" pg:"updated_at,default:now()"`
	DeletedAt    time.Time `json:"deleted_at,omitempty" pg:"deleted_at,soft_delete"`
}

type RegistrationResendValidation struct {
	Email string `json:"email" validate:"required,email"`
}

type RegistrationVerifyValidation struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

// go-pg/pg hooks
func (r *Registration) BeforeInsert(ctx context.Context) (context.Context, error) {
	r.Password = HashAndSalt(r.Password)
	return ctx, nil
}

func (r *Registration) BeforeUpdate(ctx context.Context) (context.Context, error) {
	r.Password = HashAndSalt(r.Password)
	return ctx, nil
}

// RegistrationServices represent the registration's services
type RegistrationServices interface {
	Register(context.Context, *Registration) (string, error)
	ResendCode(ctx context.Context, email string) (string, error)
	VerifyCode(ctx context.Context, email string, code string) error
}

// RegistrationRepository represent the registration's repository contract
type RegistrationRepository interface {
	StoreOrUpdateIfEmailExist(ctx context.Context, u *Registration) error
	GetByEmail(ctx context.Context, email string) (Registration, error)
	ChangeStatusToRegistered(ctx context.Context, email string) error
}
