package widget

import "fyne.io/fyne/v2/widget"

func NewEntry(placeHolder string) *widget.Entry {
	e := widget.NewEntry()
	e.SetPlaceHolder(placeHolder)

	return e
}
