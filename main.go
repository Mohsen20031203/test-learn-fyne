package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("com.example.myapp")

	// Use one of the default icons provided by fyne
	iconResource := theme.FyneLogo()

	// Set the icon for the application
	a.SetIcon(iconResource)
	fmt.Println("Icon set for the application")

	// Create a new window
	w := a.NewWindow("My Fyne App")

	// Set the icon for the window
	w.SetIcon(iconResource)
	fmt.Println("Icon set for the window")

	// Set the content of the window
	w.SetContent(container.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	// Ensure the dock icon is set on macOS
	if fyne.CurrentDevice().IsMobile() == false {
		if desk, ok := a.(interface {
			SetSystemTrayIcon(fyne.Resource)
		}); ok {
			desk.SetSystemTrayIcon(iconResource)
		}
	}

	// Set the environment variable for MacOS app bundle
	os.Setenv("FYNE_DESKTOP", "1")

	// Show and run the application
	w.ShowAndRun()
}
