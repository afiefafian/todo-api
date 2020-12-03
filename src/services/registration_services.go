package services

import (
	"context"
	"todo_api/src/entity"
)

type registrationServices struct {
	userRepo         entity.UserRepository
	registrationRepo entity.RegistrationRepository
}

// NewRegistrationServices create new register services
func NewRegistrationServices(u entity.UserRepository, r entity.RegistrationRepository) entity.RegistrationServices {
	return &registrationServices{
		userRepo:         u,
		registrationRepo: r,
	}
}

func (u *registrationServices) Register(ctx context.Context, registration *entity.Registration) error {
	//panic("implement me")
	//if err := u.registrationRepo.Store(c, user); err != nil {
	//	return err
	//}
	return nil
}

//func (u *userServices) Update(c context.Context, user *entity.User, id string) error {
//	if err := u.userRepo.Update(c, user, id); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (u *userServices) Delete(c context.Context, id string) error {
//	if err := u.userRepo.Delete(c, id); err != nil {
//		return err
//	}
//	return nil
//}
