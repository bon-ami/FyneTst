package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	tstApp := app.NewWithID("tsta")
	tstWin := tstApp.NewWindow("tstw")

	tstWin.SetContent(container.NewCenter(widget.NewButton("Click", func() {
		win := tstApp.NewWindow("tstw")
		win.SetContent(container.NewVBox(widget.NewLabel("tstt")))
		win.SetFixedSize(true)
		win.Show()
		go func(w fyne.Window) {
			time.Sleep(time.Second * 3)
			w.Close()
		}(win)
	})))
	tstWin.SetMaster()

	tstWin.Show()
	tstApp.Run()
}
