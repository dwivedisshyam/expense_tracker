package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	fw "fyne.io/fyne/v2/widget"
	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	db.New()

	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(600, 400))

	lblReg := widget.NewLabel("Registration")

	txtFname := widget.NewEntry("Firstname")
	txtLname := widget.NewEntry("Lastname")
	rGender := widget.NewRadioGroup([]string{"Male", "Female"}, nil)
	txtEmail := widget.NewEntry("Email")
	txtPswd := widget.NewPassword()

	rGender.Required = true

	form := widget.NewForm([]*fw.FormItem{
		{Text: "Firstname", Widget: txtFname},
		{Text: "Lastname", Widget: txtLname},
		{Text: "Gender", Widget: rGender},
		{Text: "Email", Widget: txtEmail},
		{Text: "Password", Widget: txtPswd},
	}, func() {
	})

	w.SetContent(container.NewVBox(
		lblReg,
		form,
	))

	w.ShowAndRun()
}
