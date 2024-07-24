package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Dark/Light Mode")

	// Function to set dark theme
	setDarkTheme := func() {
		myApp.Settings().SetTheme(theme.DarkTheme())
	}

	// Function to set light theme
	setLightTheme := func() {
		myApp.Settings().SetTheme(theme.LightTheme())
	}

	// Create buttons for dark and light mode
	darkButton := widget.NewButton("Dark Mode", func() {
		setDarkTheme()
	})

	lightButton := widget.NewButton("Light Mode", func() {
		setLightTheme()
	})

	content := container.NewVBox(
		widget.NewLabel("Hello World!"),
		darkButton,
		lightButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.ShowAndRun()
}
