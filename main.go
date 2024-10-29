package main

import (
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playMusic(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func main() {
	a := app.New()
	w := a.NewWindow("Music Player")

	var ball bool
	playButton := widget.NewButton("Play Music", func() {
		if !ball {

			go func() {
				playMusic("/Users/macbookpro/Downloads/dayone.mp3")
			}()
			ball = true
		} else {

		}
	})

	w.SetContent(playButton)
	w.Resize(fyne.NewSize(200, 100))
	w.ShowAndRun()
}
