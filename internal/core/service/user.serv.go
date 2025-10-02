package service

import (
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
	"csat-servay/pkg/logs"
)

type userServ struct {
	userRepo p.UserRepo
	raCall   p.RaCall
}

func NewUserServ(
	userRepo p.UserRepo,
	raCall p.RaCall,
) p.UserServ {
	return userServ{
		userRepo: userRepo,
		raCall:   raCall,
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

func (u userServ) CreateManyService(user []d.User) error {
	return nil
}

func (u userServ) DumpRaService() error {
	users, err := u.raCall.GetsearchAllUser()
	if err != nil {
		logs.Error(err)
		return err
	}

	if err := u.userRepo.InsertMany(users); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

// NOTE: find user repository

func (u userServ) GetAllService() (*[]d.User, error) {
	resp := []d.User{}
	if err := u.userRepo.FindAll(nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
