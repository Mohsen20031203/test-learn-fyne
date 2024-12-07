package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// شروع برنامه
	myApp := app.New()
	myWindow := myApp.NewWindow("ManageDB")

	// ایجاد ویجت‌ها
	entry := widget.NewEntry()
	entry.SetPlaceHolder("نام را وارد کنید...")

	label := widget.NewLabel("لطفاً نام خود را وارد کنید")

	button := widget.NewButton("نمایش نام", func() {
		label.SetText(fmt.Sprintf("نام وارد شده: %s", entry.Text))
	})

	// طراحی رابط کاربری
	content := container.NewVBox(
		label,
		entry,
		button,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))

	// نمایش پنجره
	myWindow.ShowAndRun()
}
