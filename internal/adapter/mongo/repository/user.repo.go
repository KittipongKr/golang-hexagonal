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
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	user *mongo.Collection
}

func NewUserRepo(user *mongo.Collection) p.UserRepo {
	return userRepo{
		user: user,
	}
}

// NOTE: insert new user repository
func (r userRepo) InsertNew(user *d.User) (*d.User, error) {
	ctx := context.Background()

	newUser := &m.User{}
	copier.Copy(newUser, user)
	newUser.SetInsertMeta()

	spew.Dump(newUser)

	if _, err := r.user.InsertOne(ctx, &newUser); err != nil {
		return nil, err
	}

	copier.Copy(user, newUser)

	return user, nil
}

func (r userRepo) InsertMany(users []d.User) error {
	ctx := context.Background()

	userModels := make([]interface{}, 0, len(users))
	for _, user := range users {
		newUser := m.User{}
		copier.Copy(&newUser, &user)
		newUser.SetInsertMeta()
		userModels = append(userModels, newUser)
	}

	if _, err := r.user.InsertMany(ctx, userModels); err != nil {
		return err
	}

	return nil
}

// NOTE: find user repository
func (r userRepo) FindAll(cond map[string]interface{}, user *[]d.User) error {
	ctx := context.Background()

	filter := bson.M{}
	if cond != nil {
		filter = h.BuildQuery(cond)
	}

	cursor, err := r.user.Find(ctx, filter)
	if err != nil {
		return err
	}

	result := []m.User{}
	if err := cursor.All(ctx, &result); err != nil {
		return err
	}

	copier.Copy(user, &result)

	return nil
}
