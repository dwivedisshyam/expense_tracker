package mongo

import (
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
)

type userStore struct {
	idGen func(time.Time) string
}

func NewUser(idGen func(time.Time) string) *userStore {
	return &userStore{
		idGen: idGen,
	}
}

func (us *userStore) Create(ctx *gofr.Context, user *model.User) (*model.User, error) {
	user.ID = us.idGen(time.Now())

	_, err := ctx.Mongo.InsertOne(ctx, CollectionUser, user)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return user, nil
}

func (us *userStore) Update(ctx *gofr.Context, user *model.User) error {
	update := bson.M{
		"$set": user,
	}

	err := ctx.Mongo.UpdateOne(ctx, CollectionUser, bson.M{"id": user.ID}, update)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *userStore) Get(ctx *gofr.Context, f *model.UserFilter) (*model.User, error) {
	var user model.User

	m := bson.M{}

	if f.ID != "" {
		m["id"] = f.ID
	}

	if f.Email != "" {
		m["email"] = f.Email
	}

	err := ctx.Mongo.FindOne(ctx, CollectionUser, m, &user)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &user, nil
}

func (us *userStore) Delete(ctx *gofr.Context, filter *model.UserFilter) error {
	_, err := ctx.Mongo.DeleteOne(ctx, CollectionUser, bson.M{"id": filter.ID})
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}
