package service

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
)

type incSvc struct {
	store store.Income
}

// func NewIncome(s store.Income) Income {
// 	return &incSvc{store: s}
// }

// func (us *incSvc) Create(i *model.Income) error {
// 	return us.store.Create(i)
// }

// func (us *incSvc) Update(i *model.Income) error {
// 	return us.store.Update(i)
// }

// func (us *incSvc) Get(i *model.Income) (*model.Income, error) {
// 	return us.store.Get(i)
// }

// func (us *incSvc) Delete(i *model.Income) error {
// 	return us.store.Delete(i)
// }
