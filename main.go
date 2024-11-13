package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Show data to byte")

	displayContainer := container.NewVBox()
	displayContainer2 := container.NewVBox()
	var file *os.File
	var Byte []byte
	var err error

	saveButton := widget.NewButton("Show data to byte", func() {

		file, err = os.Open("/Users/macbookpro/Desktop/game.jpg")
		if err != nil {
			return
		}
		defer file.Close()

		// خواندن کل محتوای فایل
		Byte, err = ioutil.ReadAll(file)

		displayBytesInChunks(Byte, displayContainer, displayContainer2)
	})

	content := container.NewVBox(
		saveButton,
		displayContainer,
	)

	m := container.NewVScroll(content)
	mainContent := container.NewBorder(nil, nil, nil, nil, m)

	columns := container.NewHSplit(mainContent, displayContainer2)
	myWindow.SetContent(columns)
	myWindow.Resize(fyne.NewSize(800, 800))
	myWindow.ShowAndRun()
}

func displayBytesInChunks(input []byte, displayContainer, displayContainer2 *fyne.Container) {
	chunkSize := 300
	textLength := len(input)

	displayContainer.Objects = nil

	for i := 0; i < textLength; i += chunkSize {
		end := i + chunkSize
		if end > textLength {
			end = textLength
		}

		buttomTest := widget.NewButton(fmt.Sprintf("part %d", i+1), func() {

			m := widget.NewMultiLineEntry()
			m.Text = fmt.Sprintf("part %d: % X", i/chunkSize+1, input[i:end])
			if len(displayContainer2.Objects) > 0 {
				displayContainer2.Objects[0] = m
			} else {
				displayContainer2.Add(m)
			}
		})

		displayContainer.Add(buttomTest)

	}
}
