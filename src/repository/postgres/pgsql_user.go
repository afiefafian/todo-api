package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"

	"todo_api/src/entity"
)

type pgsqlUserRepository struct {
	DB *pg.DB
}

// NewPostgresUserRepository create an data to represent user.Repository interface
func NewPostgresUserRepository(DB *pg.DB) entity.UserRepository {
	return &pgsqlUserRepository{DB}
}

func (p *pgsqlUserRepository) Fetch(context.Context) (res []entity.User, err error) {
	res = []entity.User{}
	err = p.DB.Model(&res).Select()
	return res, nil
}

func (p *pgsqlUserRepository) GetByID(_ context.Context, id string) (res entity.User, err error) {
	res = entity.User{}
	err = p.DB.Model(&res).Where("id = ?", id).First()
	return
}

func (p *pgsqlUserRepository) Store(_ context.Context, user *entity.User) error {
	_, err := p.DB.Model(user).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (p *pgsqlUserRepository) Update(_ context.Context, user *entity.User, id string) error {
	_, err := p.DB.Model(user).
		Column("first_name", "last_name", "email", "phone").
		Where("id = ?", id).
		Update()
	if err != nil {
		return err
	}
	return nil
}

func (p *pgsqlUserRepository) Delete(_ context.Context, id string) error {
	_, err := p.DB.Model(&entity.User{}).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (p *pgsqlUserRepository) GetByEmail(_ context.Context, email string) (res entity.User, err error) {
	res = entity.User{}
	err = p.DB.Model(&res).Where("email = ?", email).First()
	return
}
