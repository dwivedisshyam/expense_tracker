package mongo

import (
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
)

type categoryStore struct {
	idGen func(time.Time) string
}

func NewCategory(idGen func(time.Time) string) *categoryStore {
	return &categoryStore{idGen: idGen}
}

func (us *categoryStore) Index(ctx *gofr.Context, f *model.CategoryFilter) ([]model.Category, error) {
	var categories []model.Category

	err := ctx.Mongo.Find(ctx, CollectionCategory, bson.M{"user_id": f.UserID}, &categories)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return categories, nil
}

func (us *categoryStore) Create(ctx *gofr.Context, cat *model.Category) (*model.Category, error) {
	cat.ID = us.idGen(time.Now())

	_, err := ctx.Mongo.InsertOne(ctx, CollectionCategory, cat)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return cat, nil
}

func (us *categoryStore) Update(ctx *gofr.Context, cat *model.Category) error {
	update := bson.M{
		"$set": cat,
	}

	err := ctx.Mongo.UpdateOne(ctx, CollectionCategory, bson.M{"id": cat.ID, "user_id": cat.UserID}, update)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *categoryStore) Get(ctx *gofr.Context, filter *model.CategoryFilter) (*model.Category, error) {
	var category model.Category

	m := bson.M{
		"id":      filter.ID,
		"user_id": filter.UserID,
	}

	err := ctx.Mongo.FindOne(ctx, CollectionCategory, m, &category)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &category, nil
}

func (us *categoryStore) Delete(ctx *gofr.Context, cat *model.CategoryFilter) error {
	m := bson.M{
		"id":      cat.ID,
		"user_id": cat.UserID,
	}

	_, err := ctx.Mongo.DeleteOne(ctx, CollectionCategory, m)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}
