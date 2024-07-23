package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {

	myApp := app.New()

	mywindow := myApp.NewWindow("master")

	rect := canvas.NewRectangle(color.Opaque)
	mywindow.SetContent(rect)

	mywindow.Resize(fyne.NewSize(500, 500))
	mywindow.ShowAndRun()

}
