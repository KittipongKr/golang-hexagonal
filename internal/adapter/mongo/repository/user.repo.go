package repository

import (
	"context"

	h "csat-servay/internal/adapter/mongo/helper"
	m "csat-servay/internal/adapter/mongo/model"
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepo struct {
	userCol *m.MongoCollections
}

func NewUserRepo(userCol *m.MongoCollections) p.UserRepo {
	return userRepo{
		userCol: userCol,
	}
}

// NOTE: insert new user repository
func (u userRepo) InsertNew(user *d.User) (*d.User, error) {
	ctx := context.Background()

	newUser := &m.Users{}
	copier.Copy(newUser, user)
	newUser.SetInsertMeta()

	spew.Dump(newUser)

	if _, err := u.userCol.Users.InsertOne(ctx, &newUser); err != nil {
		return nil, err
	}

	copier.Copy(user, newUser)

	return user, nil
}

func (u userRepo) InsertMany(user []d.User) ([]d.User, error) {
	return nil, nil
}

// NOTE: find user repository
func (u userRepo) FindAll(cond map[string]interface{}, user *[]d.User) error {
	ctx := context.Background()

	filter := bson.M{}
	if cond != nil {
		filter = h.BuildQuery(cond)
	}

	cursor, err := u.userCol.Users.Find(ctx, filter)
	if err != nil {
		return err
	}

	result := []m.Users{}
	if err := cursor.All(ctx, &result); err != nil {
		return err
	}

	copier.Copy(user, &result)

	return nil
}
