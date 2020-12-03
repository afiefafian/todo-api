package postgres

import (
	"context"

	"github.com/go-pg/pg/v10"

	"todo_api/src/entity"
)

type pgsqlRegistrationRepository struct {
	DB *pg.DB
}

// NewPostgresRegistrationRepository create an data to represent user.Repository interface
func NewPostgresRegistrationRepository(DB *pg.DB) entity.RegistrationRepository {
	return &pgsqlRegistrationRepository{DB}
}

func (p *pgsqlRegistrationRepository) Store(ctx context.Context, register *entity.Registration) error {
	_, err := p.DB.Model(register).Insert()
	if err != nil {
		return err
	}
	return nil
}
