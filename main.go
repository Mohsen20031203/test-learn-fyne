package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/color"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hover Background Change")

	// ایجاد یک مستطیل برای پس‌زمینه
	background := canvas.NewRectangle(color.White) // رنگ پیش‌فرض پس‌زمینه
	background.SetMinSize(fyne.NewSize(400, 100))  // تنظیم اندازه مستطیل

	label := widget.NewLabel("Hover over me!")
	label.Alignment = fyne.TextAlignCenter // مرکز چینش متن

	registerBackground := canvas.NewRectangle(color.RGBA{R: 173, G: 219, B: 156, A: 200})
	registerHeader = widget.NewLabel("Registers\n16-bit words\n")
	registerHeader.TextStyle.Monospace = true
	registerHeader.TextStyle.Bold = true
	registerDisplay = cpu.GetRegisters()
	registerDisplayWidget = widget.NewLabel(registerDisplay)
	registerDisplayWidget.TextStyle.Monospace = true
	registerDisplayWidget.TextStyle.Bold = true
	registerContainer = container.NewStack(
		registerBackground,
		container.NewVBox(
			registerHeader,
			registerDisplayWidget,
		))

	// تغییر رنگ پس‌زمینه در هنگام هاور
	label.MouseIn = func() {
		background.FillColor = color.RGBA{R: 255, G: 0, B: 0, A: 255} // تغییر رنگ به قرمز
		background.Refresh()                                          // به‌روزرسانی مستطیل
	}

	label.MouseOut = func() {
		background.FillColor = color.White // برگرداندن رنگ به سفید
		background.Refresh()               // به‌روزرسانی مستطیل
	}

	// چیدمان
	content := container.New(layout.NewVBoxLayout(), background, label)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 200)) // تغییر اندازه پنجره
	myWindow.ShowAndRun()
}
