package widget

import "fyne.io/fyne/v2/widget"

func NewPassword() *widget.Entry {
	pswd := widget.NewPasswordEntry()
	pswd.PlaceHolder = `●●●●●●●●●●●●`

	return pswd
}
