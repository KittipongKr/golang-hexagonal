package port

import (
	d "csat-servay/internal/core/domain"
)

// NOTE: user repository
type UserRepo interface {
	// NOTE: insert new user repository
	InsertNew(user *d.User) (*d.User, error)
	InsertMany(user []d.User) ([]d.User, error)
	// NOTE: find user repository
	FindAll(cond map[string]interface{}, user *[]d.User) error
}

// NOTE: user service
type UserServ interface {
	// NOTE: insert new user service
	CreateNewService(user *d.User) (*d.User, error)
	CreateManyService(user []d.User) ([]d.User, error)
	// NOTE: find user repository
	GetAllService() (*[]d.User, error)
}
