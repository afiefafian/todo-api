package entity

import (
	"context"
	"time"
)

// Token entity
type Token struct {
	ID         int       `pg:"id,pk"`
	Identifier string    `pg:"identifier"`
	Type       string    `pg:"type"`
	Code       string    `pg:"code"`
	Status     bool      `pg:"status,default:true"`
	CreatedAt  time.Time `pg:"created_at,default:now()"`
	UpdatedAt  time.Time `pg:"updated_at,default:now()"`
}

// TokenRepository represent the token's repository contract
type TokenRepository interface {
	CreateNewToken(ctx context.Context, t *Token) error
	DeactivateAllByIdentifierAndType(ctx context.Context, email string, tokenType string) error
	// GetByEmail(ctx context.Context, email string) (Registration, error)
	// UpdateStatus(ctx context.Context, u *User, id string) (bool, error)
}
