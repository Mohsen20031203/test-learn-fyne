package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	// بارگذاری آیکون از فایل
	iconPath := "./icon.png"
	iconResource, err := fyne.LoadResourceFromPath(iconPath)
	if err != nil {
		fmt.Printf("خطا در بارگذاری آیکون از مسیر %s: %v\n", iconPath, err)
		return
	}

	// تنظیم آیکون برای برنامه
	a.SetIcon(iconResource)

	// ایجاد یک پنجره جدید
	w := a.NewWindow("My Fyne App")

	// تنظیم محتوای پنجره
	w.SetContent(container.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	// نمایش و اجرای برنامه
	w.ShowAndRun()
}
