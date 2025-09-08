package service

import (
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
)

type userServ struct {
	userRepo p.UserRepo
}

func NewUserServ(userRepo p.UserRepo) p.UserServ {
	return userServ{
		userRepo: userRepo,
	}
}

// NOTE: insert new user service
func (u userServ) CreateNewService(user *d.User) (*d.User, error) {
	user, err := u.userRepo.InsertNew(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userServ) CreateManyService(user []d.User) ([]d.User, error) {
	return nil, nil
}

// NOTE: find user repository

func (u userServ) GetAllService() (*[]d.User, error) {
	resp := []d.User{}
	if err := u.userRepo.FindAll(nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
