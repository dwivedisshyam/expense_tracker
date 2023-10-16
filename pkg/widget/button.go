package widget

import "fyne.io/fyne/v2/widget"

func NewButton(label string, onClick func()) *widget.Button {
	return widget.NewButton(label, onClick)
}
