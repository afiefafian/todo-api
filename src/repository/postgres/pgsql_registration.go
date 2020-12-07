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

var _ pg.BeforeInsertHook = (*entity.Registration)(nil)
var _ pg.BeforeUpdateHook = (*entity.Registration)(nil)

func (p *pgsqlRegistrationRepository) StoreOrUpdateIfEmailExist(_ context.Context, register *entity.Registration) error {
	_, err := p.DB.Model(register).
		OnConflict("(email) DO UPDATE").
		Set("first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, password = EXCLUDED.password, phone = EXCLUDED.phone").
		Insert()
	if err != nil {
		return err
	}

	return nil
}

func (p *pgsqlRegistrationRepository) GetByEmail(_ context.Context, email string) (entity.Registration, error) {
	registration := entity.Registration{}
	err := p.DB.Model(&registration).Where("email = ?", email).First()
	return registration, queryErrHandling(err)
}

func (p *pgsqlRegistrationRepository) ChangeStatusToRegistered(_ context.Context, email string) error {
	_, err := p.DB.Model(&entity.Registration{IsRegistered: true}).
		Column("is_registered").
		Where("email = ?", email).
		Update()
	if err != nil {
		return err
	}
	return nil
}
