package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Awesome TODO list")
	w.Resize(fyne.NewSize(400, 500))
	w.ShowAndRun()
}
