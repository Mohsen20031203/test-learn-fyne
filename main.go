package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("True or False Dialog")

	label := widget.NewLabel("Click the button to answer True or False.")

	// دکمه ای که دیالوگ را نشان می‌دهد
	button := widget.NewButton("Show Dialog", func() {
		// دیالوگ تایید با دو گزینه True و False
		dialog.ShowConfirm("True or False", "Please select True or False:",
			func(response bool) {
				if response {
					label.SetText("You selected: True")
					fmt.Println("User selected: True")
				} else {
					label.SetText("You selected: False")
					fmt.Println("User selected: False")
				}
			}, myWindow)
	})

	content := container.NewVBox(label, button)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
