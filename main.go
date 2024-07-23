package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myapp := app.New()
	mywindow := myapp.NewWindow("tabs")

	tabs := container.NewAppTabs(
		container.NewTabItem("mohsen", widget.NewLabel("20")),
		container.NewTabItem("ali", widget.NewLabel("25")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	mywindow.SetContent(tabs)
	mywindow.ShowAndRun()
}
