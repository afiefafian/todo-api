package entity

import (
	"context"
	"time"
)

// User entity
type User struct {
	ID        ID        `json:"id,omitempty" pg:"id,pk,default:uuid_generate_v4()"`
	FirstName string    `json:"first_name,omitempty" validate:"required" pg:"first_name"`
	LastName  string    `json:"last_name,omitempty" validate:"required" pg:"last_name"`
	Email     string    `json:"email" validate:"required,email" pg:"email"`
	Password  string    `json:"password,omitempty" pg:"password"`
	Phone     string    `json:"phone,omitempty" validate:"required" pg:"phone"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"created_at,default:now()"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"updated_at,default:now()"`
	DeletedAt time.Time `json:"deleted_at,omitempty" pg:"deleted_at,soft_delete"`
}

// UserServices represent the user's services
type UserServices interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, u *User, id string) error
	Store(context.Context, *User) error
	Delete(ctx context.Context, id string) error
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Fetch(ctx context.Context) (res []User, err error)
	GetByID(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, u *User, id string) error
	Store(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
}
