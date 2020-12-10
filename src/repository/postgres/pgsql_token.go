package postgres

import (
	"context"

	"github.com/go-pg/pg/v10"

	"github.com/afiefafian/todo-api/src/entity"
)

type pgsqlTokenRepository struct {
	DB *pg.DB
}

// NewPostgresTokenRepository create an data to represent token.Repository interface
func NewPostgresTokenRepository(DB *pg.DB) entity.TokenRepository {
	return &pgsqlTokenRepository{DB}
}

func (p pgsqlTokenRepository) CreateNewToken(_ context.Context, t *entity.Token) error {
	t.Status = true
	_, err := p.DB.Model(t).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (p pgsqlTokenRepository) DeactivateAllByIdentifierAndType(_ context.Context, email string, tokenType string) error {
	token := entity.Token{}
	_, err := p.DB.Model(&token).
		Set("status = false").
		Where("identifier = ?", email).
		Where("type = ?", tokenType).
		Update()
	if err != nil {
		return err
	}
	return nil
}
