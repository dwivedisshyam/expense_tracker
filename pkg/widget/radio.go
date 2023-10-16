package widget

import (
	"fyne.io/fyne/v2/widget"
)

func NewRadioGroup(options []string, onChange func(string)) *widget.RadioGroup {
	return widget.NewRadioGroup(options, onChange)
}
