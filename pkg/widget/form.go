package widget

import (
	"fyne.io/fyne/v2/widget"
)

func NewForm(items []*widget.FormItem, onSubmit func()) *widget.Form {
	form := &widget.Form{
		Items:    items,
		OnSubmit: onSubmit,
	}

	return form
}
