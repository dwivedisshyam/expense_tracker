package mongo

import (
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
)

type expStore struct {
	idGen func(time.Time) string
}

func NewExpense(idGen func(time.Time) string) *expStore {
	return &expStore{idGen: idGen}
}

func (us *expStore) Index(ctx *gofr.Context, f *model.ExpenseFilter) ([]model.Expense, error) {
	var exps []model.Expense

	m := bson.M{
		"user_id": f.UserID,
	}

	err := ctx.Mongo.Find(ctx, CollectionExpense, m, &exps)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return exps, nil
}

func (us *expStore) Create(ctx *gofr.Context, e *model.Expense) (*model.Expense, error) {
	e.ID = us.idGen(time.Now())

	_, err := ctx.Mongo.InsertOne(ctx, CollectionExpense, e)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return e, nil
}

func (us *expStore) Update(ctx *gofr.Context, e *model.Expense) error {
	update := bson.M{
		"$set": e,
	}

	err := ctx.Mongo.UpdateOne(ctx, CollectionExpense, bson.M{"id": e.ID, "user_id": e.UserID}, update)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *expStore) Get(ctx *gofr.Context, f *model.ExpenseFilter) (*model.Expense, error) {
	var e model.Expense

	m := bson.M{
		"id":      f.ID,
		"user_id": f.UserID,
	}

	err := ctx.Mongo.FindOne(ctx, CollectionExpense, m, &e)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &e, nil
}

func (us *expStore) Delete(ctx *gofr.Context, f *model.ExpenseFilter) error {
	m := bson.M{
		"id":      f.ID,
		"user_id": f.UserID,
	}

	_, err := ctx.Mongo.DeleteOne(ctx, CollectionExpense, m)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}
