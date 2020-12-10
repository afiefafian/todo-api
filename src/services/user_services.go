package services

import (
	"context"
	"todo_api/src/entity"
)

type userServices struct {
	userRepo entity.UserRepository
}

// NewUserServices create new user services
func NewUserServices(u entity.UserRepository) entity.UserServices {
	return &userServices{
		userRepo: u,
	}
}

func (u *userServices) Fetch(c context.Context) (res []entity.User, err error) {
	res, err = u.userRepo.Fetch(c)
	if err != nil {
		return
	}
	return
}

func (u *userServices) GetByID(c context.Context, id string) (res entity.User, err error) {
	res, err = u.userRepo.GetByID(c, id)
	if err != nil {
		return
	}
	return
}

func (u *userServices) GetByEmail(c context.Context, email string) (user entity.User, err error) {
	user, err = u.userRepo.GetByEmail(c, email)
	if err != nil {
		return
	}
	return
}

func (u *userServices) Store(c context.Context, user *entity.User) error {
	if err := u.userRepo.Store(c, user); err != nil {
		return err
	}
	return nil
}

func (u *userServices) Update(c context.Context, user *entity.User, id string) error {
	if err := u.userRepo.Update(c, user, id); err != nil {
		return err
	}
	return nil
}

func (u *userServices) Delete(c context.Context, id string) error {
	if err := u.userRepo.Delete(c, id); err != nil {
		return err
	}
	return nil
}
