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

func (p *pgsqlUserRepository) Fetch(ctx context.Context) (res []entity.User, err error) {
	res = []entity.User{}
	err = p.DB.Model(&res).Select()
	return res, nil
}

func (p *pgsqlUserRepository) GetByID(ctx context.Context, id string) (res entity.User, err error) {
	res = entity.User{}
	err = p.DB.Model(&res).Where("id = ?", id).Select()
	return
}

func (p *pgsqlUserRepository) Store(ctx context.Context, user *entity.User) error {
	_, err := p.DB.Model(user).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (p *pgsqlUserRepository) Update(ctx context.Context, user *entity.User, id string) error {
	_, err := p.DB.Model(user).
		Column("first_name, last_name, email, phone").
		Where("id = ?", id).
		Update()
	if err != nil {
		return err
	}
	return nil
}

func (p *pgsqlUserRepository) Delete(ctx context.Context, id string) error {
	_, err := p.DB.Model(&entity.User{}).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
