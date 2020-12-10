package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
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

func (p *pgsqlUserRepository) GetByID(_ context.Context, id string) (entity.User, error) {
	uid, _ := uuid.Parse(id)
	user := entity.User{
		ID: uid,
	}
	err := p.DB.Model(&user).WherePK().First()
	return user, queryErrHandling(err)
}

func (p *pgsqlUserRepository) GetByEmail(_ context.Context, email string) (entity.User, error) {
	user := entity.User{}
	err := p.DB.Model(&user).Where("email = ?", email).First()
	return user, queryErrHandling(err)
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
